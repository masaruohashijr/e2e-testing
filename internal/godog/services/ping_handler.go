package services

import (
	"fmt"
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func PingingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logging.Debug.Println("******************************** Pinging")
	logging.Debug.Println("PingingHandler")
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("SpeechResultHandler Hash: ", hash)
	fmt.Println("SpeechResultHandler TestHash: ", sTestHash)
	if hash == sTestHash {
		Ch <- "Pinged"
	}
	w.Header().Set("Status", "200 OK")
	w.Header().Set("StatusCode", "200")
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
