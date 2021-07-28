package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/ports/numbers"
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
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
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
	return nil
}

func (a *numbersAPI) AddNumber(n string) error {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/IncomingPhoneNumbers",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("PhoneNumber", n)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)

	if err != nil {
		println(err.Error())
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

func (a *numbersAPI) ListAvailableNumbers() ([]string, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/AvailablePhoneNumbers/US/%s.json",
		a.config.AccountSid, "Local")
	values := &url.Values{}

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("GET", apiEndpoint, buffer)

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
	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apnpr AvailablePhoneNumbersPageResponse
	json.Unmarshal(body, &apnpr)
	var list []string
	for _, a := range apnpr.AvailablePhoneNumbers {
		if a.PhoneNumber != "" {
			list = append(list, a.PhoneNumber)
		}
	}
	return list, nil
}

type AvailablePhoneNumbersPageResponse struct {
	AvailablePhoneNumbers []PageAvailablePhoneNumberResponse `json:"available_phone_numbers"`
	URI                   string                             `json:"uri"`
}

type PageAvailablePhoneNumberResponse struct {
	AddressRequirements string                                       `json:"address_requirements"`
	Beta                bool                                         `json:"beta"`
	Capabilities        PageAvailablePhoneNumberCapabilitiesResponse `json:"capabilities"`
	FriendlyName        string                                       `json:"friendly_name"`
	IsoCountry          string                                       `json:"iso_country"`
	Lata                *string                                      `json:"lata,omitempty"`
	Latitude            string                                       `json:"latitude"`
	Locality            *string                                      `json:"locality,omitempty"`
	Longitude           string                                       `json:"longitude"`
	PhoneNumber         string                                       `json:"phone_number"`
	PostalCode          *string                                      `json:"postal_code,omitempty"`
	RateCenter          *string                                      `json:"rate_center,omitempty"`
	Region              *string                                      `json:"region,omitempty"`
}

type PageAvailablePhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}
