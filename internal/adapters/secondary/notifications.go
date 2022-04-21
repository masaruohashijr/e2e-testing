package secondary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"zarbat_test/internal/config"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/notifications"
)

type notificationsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewNotificationsApi(config *config.ConfigType) notifications.SecondaryPort {
	return &notificationsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *notificationsAPI) ViewNotification(notificationSid string) (domains.Notification, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Notifications/%s.json", a.config.AccountSid, notificationSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummyNotification := domains.Notification{}
	if err != nil {
		return dummyNotification, err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyNotification, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	notification := domains.Notification{}
	json.Unmarshal(body, &notification)
	logging.Debug.Println(notification.Sid, notification.CallSid, notification.DateCreated, notification.Duration)
	return notification, nil
}

func (a *notificationsAPI) ListNotifications() ([]domains.Notification, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Notifications.json", a.config.AccountSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	q := req.URL.Query()
	q.Add("Page", "0")
	q.Add("PageSize", "50")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	responseNotification := domains.ResponseNotification{}
	json.Unmarshal(body, &responseNotification)
	for _, notification := range responseNotification.Notifications {
		logging.Debug.Println(notification.Sid, notification.CallSid, notification.DateCreated, notification.Duration)
	}
	return responseNotification.Notifications, nil
}
