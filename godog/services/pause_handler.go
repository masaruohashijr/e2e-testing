package services

import (
	"net/http"
	"os"
)

func PauseHandler(w http.ResponseWriter, r *http.Request) {
	println("PauseHandler")

	xml, err := os.ReadFile("xml/pause.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
