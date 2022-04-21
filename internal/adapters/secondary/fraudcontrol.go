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
	"zarbat_test/pkg/ports/fraudcontrol"
)

type fraudControlAPI struct {
	config *config.ConfigType
}

func NewFraudControlApi(config *config.ConfigType) fraudcontrol.SecondaryPort {
	return &fraudControlAPI{
		config: config,
	}
}

func (a *fraudControlAPI) BlockDestination(countryCode string) (domains.Blocked, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud/Block/%s.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("CountryCode", countryCode)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyBlocked := domains.Blocked{}
	if err != nil {
		return dummyBlocked, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyBlocked, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyBlocked, err
	}
	var blocked domains.Blocked
	err = json.Unmarshal(body, &blocked)
	if err != nil {
		return dummyBlocked, err
	}
	return blocked, nil
}

func (a *fraudControlAPI) AuthorizeDestination(countryCode string) (domains.Authorized, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud/Authorize/%s.json",
		a.config.AccountSid)
	values := &url.Values{}
	values.Add("CountryCode", countryCode)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyAuthorized := domains.Authorized{}
	if err != nil {
		return dummyAuthorized, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyAuthorized, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyAuthorized, err
	}
	var authorized domains.Authorized
	err = json.Unmarshal(body, &authorized)
	if err != nil {
		return dummyAuthorized, err
	}
	return authorized, nil
}
func (a *fraudControlAPI) ExtendDestinationAuthorization(countryCode string) (domains.Authorized, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud/Extend/%s.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("CountryCode", countryCode)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyAuthorized := domains.Authorized{}
	if err != nil {
		return dummyAuthorized, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyAuthorized, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyAuthorized, err
	}
	var extended domains.Authorized
	err = json.Unmarshal(body, &extended)
	if err != nil {
		return dummyAuthorized, err
	}
	return extended, nil
}

func (a *fraudControlAPI) WhiteListDestination(countryCode string) (domains.WhiteListed, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud/Whitelist/%s.json",
		a.config.AccountSid)
	values := &url.Values{}
	values.Add("CountryCode", countryCode)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyWhitelisted := domains.WhiteListed{}
	if err != nil {
		return dummyWhitelisted, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyWhitelisted, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyWhitelisted, err
	}
	var whitelisted domains.WhiteListed
	err = json.Unmarshal(body, &whitelisted)
	if err != nil {
		return dummyWhitelisted, err
	}
	return dummyWhitelisted, nil
}
func (a *fraudControlAPI) ListFraudControl() ([]domains.Fraud, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud.json",
		a.config.AccountSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyFrauds := []domains.Fraud{}
	if err != nil {
		return dummyFrauds, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyFrauds, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyFrauds, err
	}
	var frauds []domains.Fraud
	err = json.Unmarshal(body, &frauds)
	if err != nil {
		return dummyFrauds, err
	}
	return frauds, nil
}
