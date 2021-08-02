package services

import (
	"fmt"
	"net/http"
	"os"
	"zarbat_test/internal/logging"
)

func RecordHandler(w http.ResponseWriter, r *http.Request) {
	logging.Debug.Println("RecordHandler")
	xml, err := os.ReadFile("xml/record.xml")
	if err != nil {
		println(err.Error())
	}
	logging.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func RecordActionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rURL := r.FormValue("RecordingUrl")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("RecordingUrl")
	logging.Debug.Println(rURL)
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("RecordActionHandler Hash: ", hash)
	fmt.Println("RecordActionHandler TestHash: ", sTestHash)
	Ch <- rURL
	/*if rURL != "" { //&& hash == sTestHash {
		Ch <- rURL
	} else {
		fmt.Println("RecordActionHandler - Erro: ", rURL)
	}*/
	logging.Debug.Println("******************************** RecordAction END")
}

func TranscribeCallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, HEAD, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	transcriptionText := r.FormValue("TranscriptionText")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("Transcribe Callback")
	logging.Debug.Println("transcribed text: ", transcriptionText)
	fmt.Println("transcribed text: ", transcriptionText)
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("TranscribeCallbackHandler Hash: ", hash)
	fmt.Println("TranscribeCallbackHandler TestHash: ", sTestHash)
	if transcriptionText != "" && transcriptionText != "welcome to your new zhang account" && transcriptionText != "view in any way during your development" && transcriptionText != "let us know if we can help you in any way during your development" && transcriptionText != "can help you in any way during your development" && transcriptionText != "can i help you in any way during your development" && transcriptionText != "so we can help you in any way during your development" && transcriptionText != "help you in any way during your development" {
		if hash == sTestHash {
			fmt.Println("::: Ch <- ", transcriptionText)
			//if IsOpen(Ch) {
			Ch <- transcriptionText
			//}
			return
		}
	} else {
		fmt.Println("TranscribeCallbackHandler - Erro: ", transcriptionText)
	}
	logging.Debug.Println("******************************** Transcribe END")
	return
}

func IsOpen(ch <-chan string) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
