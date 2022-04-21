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
	"zarbat_test/pkg/ports/application"
)

type applicationAPI struct {
	config *config.ConfigType
}

func NewApplicationApi(config *config.ConfigType) application.SecondaryPort {
	return &applicationAPI{
		config: config,
	}
}

func (a *applicationAPI) UpdateApplication(applicationSid, friendlyName string) (domains.Application, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Applications/%s.json",
		a.config.AccountSid, applicationSid)

	values := &url.Values{}
	values.Add("FriendlyName", friendlyName)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyApplication := domains.Application{}
	if err != nil {
		return dummyApplication, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyApplication, err
	}
	defer resp.Body.Close()
	return dummyApplication, nil
}

func (a *applicationAPI) ListApplications() ([]domains.Application, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Applications.json",
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
	b := string(body)
	logging.Debug.Println("response Body:", b)
	responseApplication := domains.ResponseApplication{}
	json.Unmarshal(body, &responseApplication)
	for _, application := range responseApplication.Applications {
		logging.Debug.Println(application.Sid, application.Sid, application.FriendlyName, application.DateCreated)
	}
	return responseApplication.Applications, nil
}

func (a *applicationAPI) ViewApplication(applicationSid string) (domains.Application, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Applications/%s.json",
		a.config.AccountSid, applicationSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyApplication := domains.Application{}
	if err != nil {
		return dummyApplication, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyApplication, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyApplication, err
	}
	var app domains.Application
	err = json.Unmarshal(body, &app)
	if err != nil {
		return dummyApplication, err
	}
	return app, nil
}

func (a *applicationAPI) DeleteApplication(applicationSid string) (domains.Application, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Applications/%s.json",
		a.config.AccountSid, applicationSid)

	req, err := http.NewRequest("DELETE", apiEndpoint, nil)
	dummyApplication := domains.Application{}
	if err != nil {
		return dummyApplication, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyApplication, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyApplication, err
	}
	var app domains.Application
	err = json.Unmarshal(body, &app)
	if err != nil {
		return dummyApplication, err
	}
	return app, nil
}

func (a *applicationAPI) CreateApplication(friendlyName string) (domains.Application, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Applications.json",
		a.config.AccountSid)
	values := &url.Values{}
	values.Add("FriendlyName", friendlyName)
	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyApplication := domains.Application{}
	if err != nil {
		return dummyApplication, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyApplication, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyApplication, err
	}
	var app domains.Application
	err = json.Unmarshal(body, &app)
	if err != nil {
		return dummyApplication, err
	}
	return app, nil
}
