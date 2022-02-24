package services

import (
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func InboundXmlHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("InboundXmlHandler")

	//xml := "<?xml version=\"1.0\"?><Response><Play loop=\"1\">" + BaseUrl + "/mp3/sample.mp3</Play></Response>"
	//logging.Debug.Println(xml)
	xml, err := os.ReadFile("../../xml/inbound.xml")
	if err != nil {
		l.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
