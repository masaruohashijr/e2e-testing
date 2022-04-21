package main

import (
	"net/http"
	"zarbat_test/internal/logging"

	"github.com/gorilla/mux"
)

var BaseUrl = "https://mohashi.ngrok.io"

func main() {
	logging.Debug.Println("Run Server")
	r := mux.NewRouter()
	r.HandleFunc("/InboundXml", NewFunction).Methods("POST", "GET")
	r.HandleFunc("/Record", RecordHandler).Methods("POST", "GET")
	r.HandleFunc("/RecordAction", RecordActionHandler).Methods("POST", "GET")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("./mp3"))),
	)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}

func NewFunction(w http.ResponseWriter, r *http.Request) {
	response := "<?xml version=\"1.0\"?><Response><Pause length=\"10\"></Pause></Response>"
	logging.Debug.Println(response)
	w.Write([]byte(response))
}

func RecordHandler(w http.ResponseWriter, r *http.Request) {
	response := "<Response><Record background=\"true\" action=\"https://mohashi.ngrok.io/RecordAction\" maxLength=\"5\"></Record></Response>"
	logging.Debug.Println(response)
	w.Write([]byte(response))
}

func RecordActionHandler(w http.ResponseWriter, r *http.Request) {
	//RecordingUrl
	r.ParseForm()
	rURL := r.FormValue("RecordingUrl")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("RecordingUrl")
	logging.Debug.Println(rURL)
}
