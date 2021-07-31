package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	l "zarbat_test/internal/logging"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** Callback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	l.Debug.Println(fmt.Sprintf("Call Status %s.\n", status))
	if status == "completed" && CloseChannel {
		Ch <- b
	}
	l.Debug.Println("******************************** Callback END")
}

func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** FallbackHandler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	b := string(body)
	println(b)
}
