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
	"zarbat_test/pkg/ports/transcriptions"
)

type transcriptionAPI struct {
	config *config.ConfigType
}

func NewTranscriptionApi(config *config.ConfigType) transcriptions.SecondaryPort {
	return &transcriptionAPI{
		config: config,
	}
}

func (a *transcriptionAPI) ViewTranscription(transcriptionSid string) (domains.Transcription, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Transcriptions/%s.json",
		a.config.AccountSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyTranscription := domains.Transcription{}
	if err != nil {
		return dummyTranscription, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyTranscription, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyTranscription, err
	}
	var tr domains.Transcription
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return dummyTranscription, err
	}
	return tr, nil
}

func (a *transcriptionAPI) ListTranscriptions() ([]domains.Transcription, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Transcriptions.json",
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
	responseTranscription := domains.ResponseTranscription{}
	json.Unmarshal(body, &responseTranscription)
	for _, transcription := range responseTranscription.Transcriptions {
		logging.Debug.Println(transcription.Sid)
	}
	return responseTranscription.Transcriptions, nil
}

func (a *transcriptionAPI) TranscribeRecording(recordingSid string) (domains.Transcription, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Recordings/%s/Transcriptions.json",
		a.config.AccountSid, recordingSid)

	req, err := http.NewRequest("POST", apiEndpoint, nil)
	dummyTranscription := domains.Transcription{}
	if err != nil {
		return dummyTranscription, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyTranscription, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyTranscription, err
	}
	var tr domains.Transcription
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return dummyTranscription, err
	}
	return tr, nil
}

func (a *transcriptionAPI) TranscribeAudioUrl(audioUrl string) (domains.Transcription, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Transcriptions.json",
		a.config.AccountSid)
	values := &url.Values{}
	values.Add("AudioUrl", audioUrl)

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyTranscription := domains.Transcription{}
	if err != nil {
		return dummyTranscription, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyTranscription, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyTranscription, err
	}
	var tr domains.Transcription
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return dummyTranscription, err
	}
	return tr, nil
}
