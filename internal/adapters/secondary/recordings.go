package secondary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"zarbat_test/internal/config"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/recordings"
)

type recordingsAPI struct {
	config   *config.ConfigType
	VoiceUrl string `json:"VoiceUrl"`
}

func NewRecordingsApi(config *config.ConfigType) recordings.SecondaryPort {
	return &recordingsAPI{
		config:   config,
		VoiceUrl: "",
	}
}

func (a *recordingsAPI) RecordCall(callSid string, timeInSeconds int) error {
	logging.Debug.Println("=====================================================")
	logging.Debug.Println("CallSid inside RecordCall: " + callSid)
	logging.Debug.Println("=====================================================")
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Calls/%s/Recordings.json",
		a.config.AccountSid, callSid)
	values := &url.Values{}
	values.Add("Record", "true")
	values.Add("TimeLimit", strconv.Itoa(timeInSeconds))

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

func (a *recordingsAPI) ViewRecording(recordingSid string) (domains.Recording, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Recordings/%s.json", a.config.AccountSid, recordingSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummyRecording := domains.Recording{}
	if err != nil {
		return dummyRecording, err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyRecording, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	recording := domains.Recording{}
	json.Unmarshal(body, &recording)
	return recording, nil
}

func (a *recordingsAPI) DeleteRecording(recordingSid string) (domains.Recording, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Recordings/%s.json", a.config.AccountSid, recordingSid)
	req, _ := http.NewRequest("DELETE", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	encoded := EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken)
	req.Header.Add("Authorization", "Basic "+encoded)
	// TODO
	client := &http.Client{}
	resp, err := client.Do(req)
	dummyRecording := domains.Recording{}
	if err != nil {
		return dummyRecording, err
	}
	defer resp.Body.Close()
	// Print Response
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyRecording, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	recording := domains.Recording{}
	json.Unmarshal(body, &recording)
	return recording, nil
}

func (a *recordingsAPI) ListRecordings(callSid string) ([]domains.Recording, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+"/Accounts/%s/Recordings.json", a.config.AccountSid)
	req, _ := http.NewRequest("GET", apiEndpoint, nil)
	logging.Debug.Println(apiEndpoint)
	q := req.URL.Query()
	q.Add("CallSid", callSid)
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
	logging.Debug.Println("response Status:", resp.Status)
	logging.Debug.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := string(body)
	logging.Debug.Println("response Body:", b)
	responseRecording := domains.ResponseRecording{}
	json.Unmarshal(body, &responseRecording)
	for _, recording := range responseRecording.Recordings {
		logging.Debug.Println(recording.Sid, recording.CallSid, recording.DateCreated, recording.Duration)
	}
	return responseRecording.Recordings, nil
}
