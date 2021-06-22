package services

import (
	"net/http"
	"os"
)

func SayHandler(w http.ResponseWriter, r *http.Request) {
	println("SayHandler")
	xml, err := os.ReadFile("../../xml/say.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
