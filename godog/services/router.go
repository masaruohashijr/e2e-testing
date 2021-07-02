package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Ch chan string

var BaseUrl = "https://mohashi.ngrok.io"
var TestTimeout int64 = 120
var FeatureFolder = "play/play1/"
var GatherTimeOut = 60
var GatherPause = 0
var PlayPause = 3
var PlayLoop = 100
var Timeout = 10
var Background = false
var MaxLength = 10
var FileFormat = "wav"

func RunServer(c chan string) {
	Ch = c
	println("Server running")
	r := mux.NewRouter()
	r.HandleFunc("/Dial", DialHandler).Methods("POST", "GET")
	r.HandleFunc("/Ping", PingHandler).Methods("POST", "GET")
	r.HandleFunc("/Pinging", PingingHandler).Methods("POST", "GET")
	r.HandleFunc("/Pause", PauseHandler).Methods("POST", "GET")
	r.HandleFunc("/Play", PlayHandler).Methods("POST", "GET")
	r.HandleFunc("/Say", SayHandler).Methods("POST", "GET")
	r.HandleFunc("/Reject", RejectHandler).Methods("POST", "GET")
	r.HandleFunc("/Redirect", RedirectHandler).Methods("POST", "GET")
	r.HandleFunc("/RejectCallBack", RejectCallBackHandler).Methods("POST", "GET")
	r.HandleFunc("/Gather", GatherHandler).Methods("POST", "GET")
	r.HandleFunc("/Fallback", FallbackHandler).Methods("POST", "GET")
	r.HandleFunc("/Callback", CallbackHandler).Methods("POST", "GET")
	r.HandleFunc("/Hangup", HangupHandler).Methods("POST", "GET")
	r.HandleFunc("/SpeechResult", SpeechResultHandler).Methods("POST", "GET")
	r.HandleFunc("/Record", RecordHandler).Methods("POST", "GET")
	r.HandleFunc("/RecordAction", RecordActionHandler).Methods("POST", "GET")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("../../media"))),
	)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
