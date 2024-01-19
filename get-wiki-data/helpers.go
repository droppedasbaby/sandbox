package getwikidata

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func downloadFile(url string, path string, filename string) error {
	fullPath := filepath.Join(path, filename)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	length, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var downloaded int64
	buffer := make([]byte, 32*1024)
	for {
		bytesRead, err := resp.Body.Read(buffer)
		if bytesRead > 0 {
			_, writeErr := file.Write(buffer[:bytesRead])
			if writeErr != nil {
				return writeErr
			}
			downloaded += int64(bytesRead)
			showProgress(fullPath, downloaded, length)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("")
			return err
		}
	}
	fmt.Println("")
	return nil
}

func showProgress(filename string, downloaded, total int64) {
	percent := 100.0
	if total > 0 {
		percent = float64(downloaded) / float64(total) * 100
	}
	fmt.Printf("\r%s - %.2f%% downloaded", filename, percent)
}

func checkURLExists(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}
