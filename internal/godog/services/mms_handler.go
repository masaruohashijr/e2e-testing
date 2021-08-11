package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"zarbat_test/internal/logging"
)

func MmsHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("MmsHandler")
	xml, err := os.ReadFile("xml/mms.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func MmsStatusHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("*********** MMS Status Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
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
