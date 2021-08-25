package services

import (
	"fmt"
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func RecordHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("RecordHandler")
	xml, err := os.ReadFile("xml/record.xml")
	if err != nil {
		println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func RecordActionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rURL := r.FormValue("RecordingUrl")
	l.Debug.Println("************************************************")
	l.Debug.Println("RecordingUrl")
	l.Debug.Println(rURL)
	hash := r.FormValue("hash")
	sTestHash := fmt.Sprint(TestHash)
	fmt.Println("RecordActionHandler Hash: ", hash)
	l.Debug.Println("RecordActionHandler Hash: ", hash)
	fmt.Println("RecordActionHandler TestHash: ", sTestHash)
	l.Debug.Println("RecordActionHandler TestHash: ", sTestHash)
	Ch <- rURL
	/*if rURL != "" { //&& hash == sTestHash {
		Ch <- rURL
	} else {
		fmt.Println("RecordActionHandler - Erro: ", rURL)
	}*/
	l.Debug.Println("******************************** RecordAction END")
}

func TranscribeCallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, HEAD, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	transcriptionText := r.FormValue("TranscriptionText")
	l.Debug.Println("************************************************")
	l.Debug.Println("Transcribe Callback")
	l.Debug.Println("transcribed text: ", transcriptionText)
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
		l.Debug.Println("TranscribeCallbackHandler - Erro: ", transcriptionText)
	}
	l.Debug.Println("******************************** Transcribe END")
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
