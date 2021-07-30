package services

import (
	"net/http"
	"os"
)

func DialHandler(w http.ResponseWriter, r *http.Request) {
	println("DialHandler")

	xml, err := os.ReadFile("xml/dial.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
