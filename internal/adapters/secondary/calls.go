package secondary

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/ports/calls"
)

type callsAPI struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Url    string `json:"Url"`
	config *config.ConfigType
}

func NewCallsApi(config *config.ConfigType) calls.SecondaryPort {
	return &callsAPI{config: config}
}

func (a *callsAPI) MakeCall() error {
	apiEndpoint := fmt.Sprintf(a.config.GetBaseURL()+"/Accounts/%s/Calls.json", a.config.AccountSid)
	//println(apiEndpoint)
	values := &url.Values{}
	values.Add("From", a.config.From)
	values.Add("To", a.config.To)
	values.Add("Url", a.config.ActionUrl)
	values.Add("StatusCallback", a.config.StatusCallback)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("response Body:", string(body))

	return nil
}
