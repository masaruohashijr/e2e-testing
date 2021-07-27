package services

import (
	"net/http"
	"os"
)

func RecordHandler(w http.ResponseWriter, r *http.Request) {
	println("RecordHandler")
	xml, err := os.ReadFile("xml/record.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}

func RecordActionHandler(w http.ResponseWriter, r *http.Request) {
	//RecordingUrl
	r.ParseForm()
	rURL := r.FormValue("RecordingUrl")
	println("************************************************")
	println("RecordingUrl")
	println(rURL)
	Ch <- rURL
}
