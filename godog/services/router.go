package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Ch chan string

var BaseUrl = "https://e28d5a0ae640.ngrok.io"
var TestTimeout int64 = 120
var FeatureFolder = "play/play1"
var GatherTimeOut = 60
var GatherPause = 0
var PlayPause = 3
var PlayLoop = 1

func RunServer(c chan string) {
	Ch = c
	println("Server running")
	r := mux.NewRouter()
	r.HandleFunc("/Ping", PingHandler).Methods("POST")
	r.HandleFunc("/Play", PlayHandler).Methods("POST")
	r.HandleFunc("/Say", SayHandler).Methods("POST")
	r.HandleFunc("/Gather", GatherHandler).Methods("POST")
	r.HandleFunc("/Fallback", FallbackHandler).Methods("POST")
	r.HandleFunc("/Callback", CallbackHandler).Methods("POST")
	r.HandleFunc("/Pinging", PingingHandler).Methods("POST")
	r.HandleFunc("/Hangup", HangupHandler).Methods("POST")
	r.HandleFunc("/SpeechResult", SpeechResultHandler).Methods("POST", "GET")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("../../mp3"))),
	)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
