package general

import (
	"log"
	"net/http"
	"os"
)

func InboundXmlHandler(w http.ResponseWriter, r *http.Request) {
	println("InboundXmlHandler")
	xml, err := os.ReadFile("../../xml/inbound.xml")
	if err != nil {
		log.Println(err.Error)
	}
	println(string(xml))
	w.Write([]byte(xml))
}
