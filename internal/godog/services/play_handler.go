package services

import (
	"net/http"
	"os"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	println("PlayHandler")
	xml, err := os.ReadFile("xml/play.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
