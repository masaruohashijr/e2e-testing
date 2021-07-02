package services

import (
	"net/http"
	"os"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	println("Redirect Handler")
	xml, err := os.ReadFile("../../xml/redirect.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
