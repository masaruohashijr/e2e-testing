package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func InboundXmlHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("InboundXmlHandler")

	//xml := "<?xml version=\"1.0\"?><Response><Play loop=\"1\">" + BaseUrl + "/mp3/sample.mp3</Play></Response>"
	//println(xml)
	xml, err := os.ReadFile("../../xml/inbound.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
