package general

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Ch chan string

var BaseUrl = "https://4f75a27fe603.ngrok.io"

func RunServer(c chan string) {
	Ch = c
	println("Run Server")
	r := mux.NewRouter()
	r.HandleFunc("/InboundXml", InboundXmlHandler).Methods("POST")
	r.HandleFunc("/Callback", CallbackHandler).Methods("POST")
	r.HandleFunc("/Ping", PingHandler).Methods("POST")
	http.Handle("/mp3/",
		http.StripPrefix("/mp3/", http.FileServer(http.Dir("./mp3"))),
	)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}

func main() {
	c := make(chan string)
	RunServer(c)
}
