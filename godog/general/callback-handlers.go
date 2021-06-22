package general

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	r.ParseForm()                 // Parses the request body
	x := r.Form.Get("CallStatus") // x will be "" if parameter is not set
	fmt.Printf("CallStatus: %s\n", x)

	b := string(body)
	println(b)
	if strings.Contains(b, "CallStatus=completed") {
		Ch <- b
	}
}

func StatusCallbackHandler(w http.ResponseWriter, r *http.Request) {
	println("******************************** Status Callback")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println(err.Error())
	}
	b := string(body)
	println(b)
	if strings.Contains(b, "DlrStatus=delivered") {
		Ch <- b
	}
}
