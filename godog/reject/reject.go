package main

import (
	"zarbat_test/internal/config"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var NumberSecondaryPort numbers.SecondaryPort
var NumberPrimaryPort numbers.PrimaryPort
var ResponseReject domains.ResponseReject

var Ch = make(chan string)

func main() {
	println("main")
}
