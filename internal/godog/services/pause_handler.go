package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("PauseHandler")

	xml, err := os.ReadFile("xml/pause.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
