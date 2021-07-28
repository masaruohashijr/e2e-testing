package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	println("GatherHandler")
	xml, err := os.ReadFile("xml/sms.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}

func SmsStatusHanlder(w http.ResponseWriter, r *http.Request) {
	println("*********** SMS Status Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	b := string(body)
	println(b)
	if strings.Contains(b, "DlrStatus=sent") {
		Ch <- b
	}
}
