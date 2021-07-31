package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func HangupHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("HangupHandler")
	xml, err := os.ReadFile("xml/hangup.xml")
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
