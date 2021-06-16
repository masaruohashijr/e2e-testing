package main

import (
	"e2e-testing/internal/config"
	d "e2e-testing/pkg/domains"
	"e2e-testing/pkg/ports/calls"
)

var Configuration config.ConfigType
var SecondaryPort calls.SecondaryPort
var PrimaryPort calls.PrimaryPort
var ResponsePause d.ResponsePause
var Ch = make(chan int)

func main() {

}
