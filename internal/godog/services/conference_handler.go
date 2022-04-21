package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
	l "zarbat_test/internal/logging"
)

func ConferenceHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** ConferenceHandler")
	xml, err := os.ReadFile("xml/conference.xml")
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
