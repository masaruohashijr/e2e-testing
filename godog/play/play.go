package main

import (
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	"e2e-testing/pkg/ports/calls"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var ResponsePlay domains.ResponsePlay
var Ch = make(chan string)

func main() {

}
