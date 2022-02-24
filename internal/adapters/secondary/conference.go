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
	"zarbat_test/pkg/ports/conference"
)

type conferenceAPI struct {
	config *config.ConfigType
}

func NewConferenceApi(config *config.ConfigType) conference.SecondaryPort {
	return &conferenceAPI{
		config: config,
	}
}

func (a *conferenceAPI) ViewParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Conferences/%s/Participants/%s.json",
		a.config.AccountSid, conferenceSid, participantSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyParticipant := domains.Participant{}
	if err != nil {
		return dummyParticipant, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyParticipant, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyParticipant, err
	}
	var participant domains.Participant
	err = json.Unmarshal(body, &participant)
	if err != nil {
		return dummyParticipant, err
	}
	return participant, nil
}

func (a *conferenceAPI) ViewConference(conferenceSid string) (domains.Conference, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Conferences/%s.json",
		a.config.AccountSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyConference := domains.Conference{}
	if err != nil {
		return dummyConference, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyConference, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyConference, err
	}
	var conference domains.Conference
	err = json.Unmarshal(body, &conference)
	if err != nil {
		return dummyConference, err
	}
	return conference, nil
}

func (a *conferenceAPI) MuteDeafParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Conferences/%s/Participants/%s.json",
		a.config.AccountSid, conferenceSid, participantSid)
	values := &url.Values{}
	values.Add("Mute", "true")
	values.Add("Deaf", "true")

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyParticipant := domains.Participant{}
	if err != nil {
		return dummyParticipant, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyParticipant, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyParticipant, err
	}
	var participant domains.Participant
	err = json.Unmarshal(body, &participant)
	if err != nil {
		return dummyParticipant, err
	}
	return participant, nil
}
func (a *conferenceAPI) HangupParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Conferences/%s/Participants/%s.json",
		a.config.AccountSid, conferenceSid, participantSid)
	req, err := http.NewRequest("DELETE", apiEndpoint, nil)
	dummyParticipant := domains.Participant{}
	if err != nil {
		return dummyParticipant, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyParticipant, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyParticipant, err
	}
	var participant domains.Participant
	err = json.Unmarshal(body, &participant)
	if err != nil {
		return dummyParticipant, err
	}
	return participant, nil
}

func (a *conferenceAPI) PlayAudioToParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud/Whitelist/%s.json",
		a.config.AccountSid)
	values := &url.Values{}
	values.Add("AudioUrl", "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3")

	var buffer *bytes.Buffer = bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest("POST", apiEndpoint, buffer)
	dummyParticipant := domains.Participant{}
	if err != nil {
		return dummyParticipant, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyParticipant, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyParticipant, err
	}
	var participant domains.Participant
	err = json.Unmarshal(body, &participant)
	if err != nil {
		return dummyParticipant, err
	}
	return participant, nil
}

func (a *conferenceAPI) ListConferences(friendlyName string) ([]domains.Conference, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Fraud.json",
		a.config.AccountSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyConferences := []domains.Conference{}
	if err != nil {
		return dummyConferences, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyConferences, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyConferences, err
	}
	var conferences []domains.Conference
	err = json.Unmarshal(body, &conferences)
	if err != nil {
		return dummyConferences, err
	}
	return conferences, nil
}
func (a *conferenceAPI) ListParticipants(conferenceSid string) ([]domains.Participant, error) {
	apiEndpoint := fmt.Sprintf(a.config.GetApiURL()+
		"/Accounts/%s/Conferences/%s/Participants.json",
		a.config.AccountSid, conferenceSid)

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	dummyParticipants := []domains.Participant{}
	if err != nil {
		return dummyParticipants, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+EncodeToBasicAuth(a.config.AccountSid, a.config.AuthToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dummyParticipants, err
	}
	defer resp.Body.Close()
	logging.Debug.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dummyParticipants, err
	}
	var participants []domains.Participant
	err = json.Unmarshal(body, &participants)
	if err != nil {
		return dummyParticipants, err
	}
	return participants, nil
}
