package services

import (
	"net/http"
	"os"
)

func GatherHandler(w http.ResponseWriter, r *http.Request) {
	println("GatherHandler")
	xml, err := os.ReadFile("xml/gather.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}

func SpeechResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sr := r.FormValue("SpeechResult")
	println("************************************************")
	println("SpeechResult")
	println(sr)
	if sr == "WORKED" {
		println("Testing ... continue")
	} else {
		Ch <- sr
	}
}
