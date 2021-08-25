package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func SayHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("SayHandler")
	xml, err := os.ReadFile("xml/say.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
