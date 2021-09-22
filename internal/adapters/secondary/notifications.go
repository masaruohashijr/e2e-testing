package secondary

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/sms"
)

type notificationsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewNotificationsApi(config *config.ConfigType) sms.SecondaryPort {
	return &smsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *smsAPI) SendNotification(to, from, message string) error {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/SMS/Messages.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("To", to)
	values.Add("From", from)
	values.Add("Body", message)
	values.Add("StatusCallback", services.BaseUrl+"/SmsStatus")

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *smsAPI) ViewNotification(smsSid string) (domains.Sms, error) {
	s := domains.Sms{}
	return s, nil
}

func (a *smsAPI) ListNotifications(from, to string) ([]domains.Sms, error) {
	s := domains.Sms{}
	ss := []domains.Sms{s}
	return ss, nil
}
