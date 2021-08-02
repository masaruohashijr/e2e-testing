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
	println(apiEndpoint)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	println("Basic " + encoded)
	req.Header.Add("Authorization", "Basic "+encoded)
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
	b := string(body)
	fmt.Println("response Body:", b)
	c := Call{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		println(err.Error())
	}
	println("CallSid: ", c.Sid)
	services.CallSidContext = c.Sid
	return nil
}

type Call struct {
	DateUpdated     string `json:"date_updated,omitempty"`
	ParentCallSid   string `json:"parent_call_sid,omitempty"`
	Duration        int    `json:"duration,omitempty"`
	From            string `json:"from,omitempty"`
	To              string `json:"to,omitempty"`
	CallerIdBlocked string `json:"caller_id_blocked,omitempty"`
	AnsweredBy      string `json:"answered_by,omitempty"`
	Sid             string `json:"sid,omitempty"`
	RecordingsCount string `json:"recordings_count,omitempty"`
	Price           string `json:"price,omitempty"`
	ApiVersion      string `json:"api_version,omitempty"`
	Status          string `json:"status,omitempty"`
	Direction       string `json:"direction,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	DateCreated     string `json:"date_created,omitempty"`
	ForwardedFrom   string `json:"forwarded_from,omitempty"`
	AccountSid      string `json:"account_sid,omitempty"`
	DurationBilled  int    `json:"duration_billed,omitempty"`
	PhoneNumberSid  string `json:"phone_number_sid,omitempty"`
}
