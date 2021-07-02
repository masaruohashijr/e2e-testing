package main

import (
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	"e2e-testing/pkg/ports/calls"
	"e2e-testing/pkg/ports/numbers"
)

var Configuration config.ConfigType
var ResponsePing domains.ResponsePing
var CallsSecondaryPort calls.SecondaryPort
var CallsPrimaryPort calls.PrimaryPort
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var Ch = make(chan string)

func main() {

}
