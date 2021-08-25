package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func ConferenceHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** ConferenceHandler")
	xml, err := os.ReadFile("xml/conference.xml")
	if err != nil {
		println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
