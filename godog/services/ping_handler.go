package services

import (
	"net/http"
	"os"
)

func PingingHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Ping")
	Ch <- "Pinged"
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	println("PingHandler")

	xml, err := os.ReadFile("xml/ping.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
