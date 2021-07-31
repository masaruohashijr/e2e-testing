package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func GatherHandler(w http.ResponseWriter, r *http.Request) {
	//logging.Debug.Println("GatherHandler")
	xml, err := os.ReadFile("xml/gather.xml")
	if err != nil {
		println(err.Error())
	}
	//logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func SpeechResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sr := r.FormValue("SpeechResult")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("SpeechResult")
	logging.Debug.Println(sr)
	if sr == "WORKED" {
		logging.Debug.Println("Testing ... continue")
	} else {
		Ch <- sr
	}
}
