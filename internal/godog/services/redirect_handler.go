package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("Redirect Handler")
	xml, err := os.ReadFile("xml/redirect.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
