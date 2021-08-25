package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	l "zarbat_test/internal/logging"
)

func RejectHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("Reject Handler")
	xml, err := os.ReadFile("../../xml/reject.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func RejectCallBackHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** RejectCallback START")
	fmt.Println("******************************** RejectCallback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Debug.Println(err.Error())
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	fmt.Println(b)
	l.Debug.Println(b)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	callSid := url_parameters["CallSid"][0]
	l.Debug.Println(fmt.Sprintf("Call Status %s.\n", status))
	fmt.Printf("Call Status %s.\n", status)
	l.Debug.Println(fmt.Sprintf("Call Sid %s.\n", callSid))
	fmt.Printf("Call Sid %s.\n", callSid)
	l.Debug.Println(fmt.Sprintf("Call Sid Context %s.\n", CallSidContext))
	fmt.Printf("Call Sid Context %s.\n", CallSidContext)
	if status == "canceled" && callSid == CallSidContext {
		Ch <- "Call Cancelled"
	}
	l.Debug.Println("******************************** RejectCallback END")
	fmt.Println("******************************** RejectCallback END")
}
