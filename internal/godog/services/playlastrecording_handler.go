package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func PlayLastRecordingHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("PlayLastRecordingHandler")
	xml, err := os.ReadFile("xml/playlastrecording.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
