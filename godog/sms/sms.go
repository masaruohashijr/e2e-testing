package main

import (
	"zarbat_test/internal/config"
	d "zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
)

var Configuration config.ConfigType
var SecondaryPort calls.SecondaryPort
var PrimaryPort calls.PrimaryPort
var ResponseSMS d.ResponseSMS
var Ch = make(chan string)

func main() {

}
