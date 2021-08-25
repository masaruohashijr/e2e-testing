package services

import (
	"fmt"
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func PingingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	l.Debug.Println("******************************** Pinging")
	l.Debug.Println("PingingHandler")
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("SpeechResultHandler Hash: ", hash)
	l.Debug.Println("SpeechResultHandler Hash: ", hash)
	fmt.Println("SpeechResultHandler TestHash: ", sTestHash)
	l.Debug.Println("SpeechResultHandler TestHash: ", sTestHash)
	if hash == sTestHash {
		Ch <- "Pinged"
	}
	w.Header().Set("Status", "200 OK")
	w.Header().Set("StatusCode", "200")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("PingHandler")
	xml, err := os.ReadFile("xml/ping.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
