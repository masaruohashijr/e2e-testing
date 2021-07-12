package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Callback START")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()
	b := string(body)
	url_parameters, err := url.ParseQuery(b)
	status := url_parameters["CallStatus"][0]
	fmt.Printf("Call Status %s.\n", status)
	if status == "completed" {
		Ch <- b
	}
	println("******************************** Callback END")
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
