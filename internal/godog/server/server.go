package main

import (
	"fmt"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
)

func main() {
	Ch := make(chan string)
	services.RunServer(Ch, false)
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		Ch = nil
	}
	return
}
