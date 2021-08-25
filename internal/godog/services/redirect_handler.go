package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("Redirect Handler")
	xml, err := os.ReadFile("xml/redirect.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
