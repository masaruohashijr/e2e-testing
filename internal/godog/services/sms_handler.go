package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"zarbat_test/internal/logging"
	l "zarbat_test/internal/logging"
)

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("SmsHandler")
	xml, err := os.ReadFile("xml/sms.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func SmsStatusHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("*********** SMS Status Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logging.Debug.Println(err.Error())
		l.Debug.Println(err.Error())
	}
	bodyContent := string(body)
	url_parameters, _ := url.ParseQuery(bodyContent)
	if url_parameters["DlrStatus"] != nil {
		status := url_parameters["DlrStatus"][0]
		if status == "sent" {
			Ch <- bodyContent
		}
	}
	w.Header().Set("Allow", "GET, HEAD, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
	return
}
