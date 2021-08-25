package services

import (
	"fmt"
	"net/http"
	"os"
	"zarbat_test/internal/logging"
	l "zarbat_test/internal/logging"
)

func GatherHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("GatherHandler")
	xml, err := os.ReadFile("xml/gather.xml")
	if err != nil {
		println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func SpeechResultHandler(w http.ResponseWriter, r *http.Request) {
	/* 	if _, ok := <-Ch; !ok {
		return
	} */
	r.ParseForm()
	sr := r.FormValue("SpeechResult")
	dg := r.FormValue("Digits")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("SpeechResult")
	logging.Debug.Println(sr)
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("SpeechResultHandler Hash: ", hash)
	fmt.Println("SpeechResultHandler TestHash: ", sTestHash)
	fmt.Println("SpeechResult: ", sr)
	fmt.Println("Digits: ", dg)
	if (sr != "" && sr != "welcome to your new Zhang account" || dg != "") && hash == sTestHash {
		//f IsOpen(Ch) {
		Ch <- sr
		//}
	}
	w.Header().Set("Allow", "GET, HEAD, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
	return
}
