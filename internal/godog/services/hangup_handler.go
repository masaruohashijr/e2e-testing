package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func HangupHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("HangupHandler")
	xml, err := os.ReadFile("xml/hangup.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
