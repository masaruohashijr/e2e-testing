package general

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(c chan int) {
	Ch = c
	println("Run Server")
	r := mux.NewRouter()
	r.HandleFunc("/InboundXml", InboundXmlHandler).Methods("POST")
	r.HandleFunc("/Callback", CallbackhHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
