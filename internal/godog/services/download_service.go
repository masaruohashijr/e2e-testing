package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	l "zarbat_test/internal/logging"
)

func DownloadFile(filepath string, url string) (err error) {
	l.Debug.Println("******************************** DownloadFile")
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
