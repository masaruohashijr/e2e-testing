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
	"zarbat_test/pkg/ports/carrier"
)

type carrierAPI struct {
	config *config.ConfigType
}

func NewCarrierApi(config *config.ConfigType) carrier.SecondaryPort {
	return &carrierAPI{
		config: config,
	}
}

func (a *carrierAPI) CarrierLookup(phoneNumber string) (domains.CarrierLookup, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Carrier.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("PhoneNumber", phoneNumber)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyCarrierLookup := domains.CarrierLookup{}
	if err != nil {
		return dummyCarrierLookup, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyCarrierLookup, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyCarrierLookup, err
	}
	var carrierLookup domains.CarrierLookup
	err = json.Unmarshal(body, &carrierLookup)
	if err != nil {
		return dummyCarrierLookup, err
	}
	return carrierLookup, nil
}

func (a *carrierAPI) CarrierLookupList() ([]domains.CarrierLookup, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Carrier.json",
		a.config.AccountSid)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyCarrierLookups := []domains.CarrierLookup{}
	if err != nil {
		return dummyCarrierLookups, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyCarrierLookups, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyCarrierLookups, err
	}
	var responseCarrierLookups domains.ResponseCarrierLookup
	err = json.Unmarshal(body, &responseCarrierLookups)
	if err != nil {
		return dummyCarrierLookups, err
	}
	return responseCarrierLookups.CarrierLookups, nil
}
func (a *carrierAPI) CNAMLookup(phoneNumber string) (domains.CNAM, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Cnam.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("PhoneNumber", phoneNumber)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyCNAM := domains.CNAM{}
	if err != nil {
		return dummyCNAM, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyCNAM, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyCNAM, err
	}
	var cnam domains.CNAM
	err = json.Unmarshal(body, &cnam)
	if err != nil {
		return dummyCNAM, err
	}
	return cnam, nil
}

func (a *carrierAPI) CNAMLookupList() ([]domains.CNAM, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Cnam.json",
		a.config.AccountSid)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyCNAM := []domains.CNAM{}
	if err != nil {
		return dummyCNAM, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyCNAM, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyCNAM, err
	}
	var responseCNAM domains.ResponseCNAM
	err = json.Unmarshal(body, &responseCNAM)
	if err != nil {
		return dummyCNAM, err
	}
	return responseCNAM.CnamLookups, nil
}
func (a *carrierAPI) BNALookup(phoneNumber string) (domains.BNA, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Bna.json",
		a.config.AccountSid)

	values := &url.Values{}
	values.Add("PhoneNumber", phoneNumber)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyBNA := domains.BNA{}
	if err != nil {
		return dummyBNA, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyBNA, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyBNA, err
	}
	var bnaLookup domains.BNA
	err = json.Unmarshal(body, &bnaLookup)
	if err != nil {
		return dummyBNA, err
	}
	return bnaLookup, nil
}

func (a *carrierAPI) BNALookupList() ([]domains.BNA, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Lookups/Bna.json",
		a.config.AccountSid)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyBNALookups := []domains.BNA{}
	if err != nil {
		return dummyBNALookups, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyBNALookups, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyBNALookups, err
	}
	var responseBNALookups domains.ResponseBNA
	err = json.Unmarshal(body, &responseBNALookups)
	if err != nil {
		return dummyBNALookups, err
	}
	return responseBNALookups.BNALookups, nil
}
