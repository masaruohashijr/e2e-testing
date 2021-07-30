package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func RejectHandler(w http.ResponseWriter, r *http.Request) {
	println("Reject Handler")
	xml, err := os.ReadFile("../../xml/reject.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}

func RejectCallBackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Reject Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}

	b := string(body)
	println(b)
	if strings.Contains(b, "CallStatus=canceled") {
		Ch <- "Call Cancelled"
	}
}
