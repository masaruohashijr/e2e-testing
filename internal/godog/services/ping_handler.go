package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func PingingHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("******************************** Ping")
	Ch <- "Pinged"
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("PingHandler")

	xml, err := os.ReadFile("xml/ping.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
