package services

import (
	"net/http"
	"zarbat_test/internal/logging"

	"github.com/gorilla/mux"
)

var Ch chan string
var CloseChannel bool
var ChComplete chan string

var BaseUrl string
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
var TestHash uint32

func RunServer(c chan string, close bool) {
	Ch = c
	println("RENEWING CHANNEL")
	CloseChannel = close
	if router != nil {
		return
	}
	if logging.Debug != nil {
		logging.Debug.Println("Server running")
	}
	router = mux.NewRouter()
	router.HandleFunc("/Callback", CallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Conference", ConferenceHandler).Methods("POST", "GET")
	router.HandleFunc("/ConferenceCallback", ConferenceCallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Dial", DialHandler).Methods("POST", "GET")
	router.HandleFunc("/DialCallback", DialCallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Fallback", FallbackHandler).Methods("POST", "GET")
	router.HandleFunc("/Gather", GatherHandler).Methods("POST", "GET")
	router.HandleFunc("/Hangup", HangupHandler).Methods("POST", "GET")
	router.HandleFunc("/Media", MediaHandler).Methods("POST", "GET")
	router.HandleFunc("/Mms", MmsHandler).Methods("POST", "GET")
	router.HandleFunc("/MmsStatus", MmsStatusHandler).Methods("POST", "GET")
	router.HandleFunc("/Number", NumberHandler).Methods("POST", "GET")
	router.HandleFunc("/Pause", PauseHandler).Methods("POST", "GET")
	router.HandleFunc("/Ping", PingHandler).Methods("POST", "GET")
	router.HandleFunc("/Pinging", PingingHandler).Methods("POST", "GET")
	router.HandleFunc("/Play", PlayHandler).Methods("POST", "GET")
	router.HandleFunc("/PlayLastRecording", PlayLastRecordingHandler).Methods("POST", "GET")
	router.HandleFunc("/Record", RecordHandler).Methods("POST", "GET")
	router.HandleFunc("/RecordAction", RecordActionHandler).Methods("POST", "GET")
	router.HandleFunc("/Redirect", RedirectHandler).Methods("POST", "GET")
	router.HandleFunc("/Reject", RejectHandler).Methods("POST", "GET")
	router.HandleFunc("/RejectCallBack", RejectCallBackHandler).Methods("POST", "GET")
	router.HandleFunc("/Say", SayHandler).Methods("POST", "GET")
	router.HandleFunc("/Sms", SmsHandler).Methods("POST", "GET")
	router.HandleFunc("/SmsStatus", SmsStatusHandler).Methods("POST", "GET")
	router.HandleFunc("/SpeechResult", SpeechResultHandler).Methods("POST", "GET")
	router.HandleFunc("/TranscribeCallback", TranscribeCallbackHandler).Methods("POST", "GET")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("../../media"))),
	)
	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
