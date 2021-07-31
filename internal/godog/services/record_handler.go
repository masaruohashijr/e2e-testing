package services

import (
	"fmt"
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func RecordHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("RecordHandler")
	xml, err := os.ReadFile("xml/record.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func RecordActionHandler(w http.ResponseWriter, r *http.Request) {
	//RecordingUrl
	r.ParseForm()
	rURL := r.FormValue("RecordingUrl")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("RecordingUrl")
	logging.Debug.Println(rURL)
	Ch <- rURL
	logging.Debug.Println("******************************** RecordAction END")
}

func TranscribeCallbackHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rURL := r.FormValue("TranscriptionText")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("Transcribe Callback")
	logging.Debug.Println("transcribed text: ", rURL)
	fmt.Println("transcribed text: ", rURL)
	if rURL != "" {
		Ch <- rURL
	} else {
		fmt.Println("TranscribeCallbackHandler - Erro: ", rURL)
	}
	logging.Debug.Println("******************************** Transcribe END")
}
