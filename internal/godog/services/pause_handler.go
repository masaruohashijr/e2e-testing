package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("PauseHandler")

	xml, err := os.ReadFile("xml/pause.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
