package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("PlayHandler")
	xml, err := os.ReadFile("xml/play.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
