package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/sms"
)

type smsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewSmsApi(config *config.ConfigType) sms.SecondaryPort {
	return &smsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *smsAPI) SendSMS(to, from, message string) error {
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

func (a *smsAPI) ViewSMS(smsSid string) (domains.Sms, error) {
	s := domains.Sms{}
	return s, nil
}

func (a *smsAPI) ListSMS(from, to string) ([]domains.Sms, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/SMS/Messages", a.config.AccountSid)
	values := &url.Values{}
	values.Add("From", a.config.From)
	values.Add("To", a.config.To)
	values.Add("DateSent", "<="+time.Now().Format("2021-09-21"))
	values.Add("Page", "0")
	values.Add("PageSize", "1")
	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("GET", apiEndpoint, buffer)
	println(apiEndpoint)

	if err != nil {
		return nil, err
	}

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
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	smsResponse := domains.SMSResponse{}
	json.Unmarshal(body, &smsResponse)
	for _, sms := range smsResponse.SmsMessages {
		fmt.Println(sms.From, sms.To, sms.DateCreated, sms.Body)
	}
	return smsResponse.SmsMessages, nil
}
