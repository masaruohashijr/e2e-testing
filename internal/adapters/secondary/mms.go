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
	"zarbat_test/pkg/ports/mms"
)

type mmsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewMmsApi(config *config.ConfigType) mms.SecondaryPort {
	return &mmsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *mmsAPI) SendMMS(from, to, message string) error {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/MMS/Messages.json",
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

func (a *mmsAPI) ViewMMS(mmsSid string) (domains.Mms, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/MMS/Messages/%s.json", a.config.AccountSid, mmsSid)
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
	dummyMms := domains.Mms{}
	if err != nil {
		return dummyMms, err
	}
	defer resp.Body.Close()
	// Print Response
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyMms, err
	}
	b := string(body)
	fmt.Println("response Body:", b)
	mms := domains.Mms{}
	json.Unmarshal(body, &mms)
	return mms, nil
}

func (a *mmsAPI) ListMMS(from, to string) ([]domains.Mms, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/MMS/Messages.json", a.config.AccountSid)
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
	mmsResponse := domains.MMSResponse{}
	json.Unmarshal(body, &mmsResponse)
	for _, mms := range mmsResponse.MmsMessages {
		fmt.Println(mms.From, mms.To, mms.DateCreated, mms.Body)
	}
	return mmsResponse.MmsMessages, nil
}
