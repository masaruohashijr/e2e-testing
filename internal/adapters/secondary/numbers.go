package secondary

import (
	"bytes"
	"e2e-testing/internal/config"
	"e2e-testing/pkg/ports/numbers"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type numbersAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewNumbersApi(config *config.ConfigType) numbers.SecondaryPort {
	return &numbersAPI{
		VoiceUrl: "",
		config:   config,
	}
}

func (a *numbersAPI) UpdateNumber() error {
	apiEndpoint := fmt.Sprintf(a.config.GetBaseURL()+
		"/Accounts/%s/IncomingPhoneNumbers/%s.json",
		a.config.AccountSid, a.config.ToSid)

	values := &url.Values{}
	values.Add("VoiceUrl", a.config.VoiceUrl)

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
	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("response Body:", string(body))

	return nil
}
