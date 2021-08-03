package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	l "zarbat_test/internal/logging"
)

var CallSidContext = ""

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** Callback START")
	fmt.Println("******************************** Callback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	fmt.Println(b)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	callSid := url_parameters["CallSid"][0]
	l.Debug.Println(fmt.Sprintf("Call Status %s.\n", status))
	fmt.Printf("Call Status %s.\n", status)
	fmt.Printf("Call Sid %s.\n", callSid)
	fmt.Printf("Call Sid Context %s.\n", CallSidContext)
	if status == "completed" && CloseChannel && callSid == CallSidContext {
		Ch <- b
	}
	l.Debug.Println("******************************** Callback END")
	fmt.Println("******************************** Callback END")
}

func DialCallbackHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** Dial Callback START")
	fmt.Println("******************************** Dial Callback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	fmt.Println(b)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	callSid := url_parameters["CallSid"][0]
	l.Debug.Println(fmt.Sprintf("Call Status %s.\n", status))
	fmt.Printf("Call Status %s.\n", status)
	fmt.Printf("Call Sid %s.\n", callSid)
	fmt.Printf("Call Sid Context %s.\n", CallSidContext)
	if status == "completed" && CloseChannel {
		Ch <- b
	}
	l.Debug.Println("******************************** Dial Callback END")
	fmt.Println("******************************** Dial Callback END")
}

func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** FallbackHandler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	b := string(body)
	println(b)
}

func ConferenceCallbackHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** Callback START")
	fmt.Println("******************************** Callback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	fmt.Println(b)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	callSid := url_parameters["CallSid"][0]
	l.Debug.Println(fmt.Sprintf("Call Status %s.\n", status))
	fmt.Printf("Call Status %s.\n", status)
	fmt.Printf("Call Sid %s.\n", callSid)
	fmt.Printf("Call Sid Context %s.\n", CallSidContext)
	if url_parameters["ConferenceName"] != nil {
		conferenceName := url_parameters["ConferenceName"][0]
		if status == "completed" && conferenceName != "" {
			Ch <- b
		}
	}
	l.Debug.Println("******************************** Callback END")
	fmt.Println("******************************** Callback END")
}
