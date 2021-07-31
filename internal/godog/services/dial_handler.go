package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func DialHandler(w http.ResponseWriter, r *http.Request) {
	println("DialHandler")

	xml, err := os.ReadFile("xml/dial.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
