package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/account"
)

type accountAPI struct {
	config       *config.ConfigType
	FriendlyName string `json:"friendly_name"`
}

func NewAccountApi(config *config.ConfigType) account.SecondaryPort {
	return &accountAPI{
		FriendlyName: "",
		config:       config,
	}
}

func (a *accountAPI) UpdateAccount(friendlyName string) error {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("FriendlyName", friendlyName)

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

func (a *accountAPI) ViewAccount() (*domains.Account, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s.json",
		a.config.AccountSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var acc domains.Account
	err = json.Unmarshal(body, &acc)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
