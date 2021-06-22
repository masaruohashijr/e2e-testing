package services

import (
	"net/http"
	"os"
)

func HangupHandler(w http.ResponseWriter, r *http.Request) {
	println("HangupHandler")
	xml, err := os.ReadFile("../../xml/hangup.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
