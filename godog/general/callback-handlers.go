package general

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var Ch chan int

func CallbackhHandler(w http.ResponseWriter, r *http.Request) {
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
		data := strings.Split(b, "&")
		duration := data[3]
		d, _ := strconv.Atoi(duration)
		fmt.Printf("Duration %d\n", duration)
		Ch <- d
	}
}
