package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
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
	values.Add("From", from)
	values.Add("To", to)
	values.Add("Body", message)

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
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	defer resp.Body.Close()
	return nil
}

func (a *smsAPI) ViewSMS(smsSid string) (domains.Sms, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/SMS/Messages/%s.json", a.config.AccountSid, smsSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	println(apiEndpoint)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummySms := domains.Sms{}
	if err != nil {
		return dummySms, err
	}
	defer resp.Body.Close()
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummySms, err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	sms := domains.Sms{}
	json.Unmarshal(body, &sms)
	return sms, nil
}

func (a *smsAPI) ListSMS(from, to string) ([]domains.Sms, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/SMS/Messages.json", a.config.AccountSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	println(apiEndpoint)
	q := req.URL.Query()
	q.Add("From", from)
	q.Add("To", to)
	q.Add("Page", "0")
	q.Add("PageSize", "1")
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
