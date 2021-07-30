package services

import (
	"net/http"
	"os"
)

func InboundXmlHandler(w http.ResponseWriter, r *http.Request) {
	println("InboundXmlHandler")

	//xml := "<?xml version=\"1.0\"?><Response><Play loop=\"1\">" + BaseUrl + "/mp3/sample.mp3</Play></Response>"
	//println(xml)
	xml, err := os.ReadFile("../../xml/inbound.xml")
	if err != nil {
		println(err.Error())
	}
	println(string(xml))
	w.Write([]byte(xml))
}
