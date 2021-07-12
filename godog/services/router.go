package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Ch chan string

var BaseUrl = "http://mohashi.ngrok.io"
var TestTimeout int64 = 60
var GatherTimeOut = 60
var GatherPause = 0
var PlayPause = 3
var PlayLoop = 100
var Timeout = 10
var Background = false
var MaxLength = 10
var FileFormat = "wav"
var router *mux.Router

func RunServer(c chan string) {
	Ch = c
	if router != nil {
		return
	}
	println("Server running")
	router = mux.NewRouter()
	router.HandleFunc("/Dial", DialHandler).Methods("POST", "GET")
	router.HandleFunc("/Ping", PingHandler).Methods("POST", "GET")
	router.HandleFunc("/Pinging", PingingHandler).Methods("POST", "GET")
	router.HandleFunc("/Pause", PauseHandler).Methods("POST", "GET")
	router.HandleFunc("/Play", PlayHandler).Methods("POST", "GET")
	router.HandleFunc("/Say", SayHandler).Methods("POST", "GET")
	router.HandleFunc("/Reject", RejectHandler).Methods("POST", "GET")
	router.HandleFunc("/Redirect", RedirectHandler).Methods("POST", "GET")
	router.HandleFunc("/RejectCallBack", RejectCallBackHandler).Methods("POST", "GET")
	router.HandleFunc("/Gather", GatherHandler).Methods("POST", "GET")
	router.HandleFunc("/Fallback", FallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Callback", CallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Hangup", HangupHandler).Methods("POST", "GET")
	router.HandleFunc("/SpeechResult", SpeechResultHandler).Methods("POST", "GET")
	router.HandleFunc("/sms", SmsHandler).Methods("POST", "GET")
	router.HandleFunc("/SmsStatus", SmsStatusHanlder).Methods("POST", "GET")
	router.HandleFunc("/Record", RecordHandler).Methods("POST", "GET")
	router.HandleFunc("/RecordAction", RecordActionHandler).Methods("POST", "GET")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("../../media"))),
	)
	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
