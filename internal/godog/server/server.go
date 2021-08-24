package main

import (
	"fmt"
	"time"
	"zarbat_test/internal/godog/services"
)

func main() {
	Ch := make(chan string)
	services.RunServer(Ch, false)
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
	}
	return
}
