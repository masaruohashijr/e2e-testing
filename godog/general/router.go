package general

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Ch chan string

func RunServer(c chan string) {
	Ch = c
	println("Run Server")
	r := mux.NewRouter()
	r.HandleFunc("/InboundXml", InboundXmlHandler).Methods("POST")
	r.HandleFunc("/Callback", CallbackHandler).Methods("POST")
	r.HandleFunc("/StatusCallback", StatusCallbackHandler).Methods("POST")
	r.HandleFunc("/Ping", PingHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
