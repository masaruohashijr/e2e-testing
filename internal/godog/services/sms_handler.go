package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"zarbat_test/internal/logging"
)

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("GatherHandler")
	xml, err := os.ReadFile("xml/sms.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func SmsStatusHanlder(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("*********** SMS Status Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	b := string(body)
	logging.Debug.Println(b)
	if strings.Contains(b, "DlrStatus=sent") {
		Ch <- b
	}
}
