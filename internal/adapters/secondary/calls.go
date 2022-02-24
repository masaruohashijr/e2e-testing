package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
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
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Calls.json", a.config.AccountSid)
	values := &url.Values{}
	values.Add("From", a.config.From)
	values.Add("To", a.config.To)
	values.Add("Url", a.config.ActionUrl)
	values.Add("StatusCallback", a.config.StatusCallback)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	logging.Debug.Println(apiEndpoint)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	logging.Debug.Println("Basic " + encoded)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	c := domains.Call{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	logging.Debug.Println("CallSid: ", c.Sid)
	services.CallSidContext = c.Sid
	return nil
}

func (a *callsAPI) FilterCalls(from, to, status string) ([]domains.Call, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Calls.json", a.config.AccountSid)
	logging.Debug.Println(apiEndpoint)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	q := req.URL.Query()
	q.Add("From", from)
	q.Add("To", to)
	q.Add("Status", status)
	req.URL.RawQuery = q.Encode()

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
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	callsResponse := domains.CallsResponse{}
	json.Unmarshal(body, &callsResponse)
	for _, c := range callsResponse.Calls {
		logging.Debug.Println(c.From, c.To, c.DateCreated, c.Duration)
	}
	return callsResponse.Calls, nil
}

func (a *callsAPI) ListCalls() ([]domains.Call, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Calls.json", a.config.AccountSid)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)

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
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	callsResponse := domains.CallsResponse{}
	json.Unmarshal(body, &callsResponse)
	for _, c := range callsResponse.Calls {
		logging.Debug.Println(c.From, c.To, c.DateCreated, c.Duration)
	}
	return callsResponse.Calls, nil
}

func (a *callsAPI) ViewCall(callSid string) (domains.Call, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Calls/%s.json", a.config.AccountSid, callSid)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummyCall := domains.Call{}
	if err != nil {
		return dummyCall, err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyCall, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	c := domains.Call{}
	json.Unmarshal(body, &c)
	logging.Debug.Println(c.From, c.To, c.DateCreated, c.Duration)
	return c, nil
}
