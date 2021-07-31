package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("PlayHandler")
	xml, err := os.ReadFile("xml/play.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
