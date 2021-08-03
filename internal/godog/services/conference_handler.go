package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func ConferenceHandler(w http.ResponseWriter, r *http.Request) {
	println("ConferenceHandler")

	xml, err := os.ReadFile("xml/conference.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
