package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var targetDir = "."
var serverUrl = ""

func main() {
	fmt.Println("Batch OCR Request Demo Utility V1.0")
	if len(os.Args) > 2 {
		serverUrl = os.Args[1]
		if !strings.HasPrefix(serverUrl, "http://") {
			serverUrl = "http://" + serverUrl
		}
		fmt.Println("Processing file: ", os.Args[2])
		targetDir = os.Args[2]
	} else {
		fmt.Println("Usage: batchocr <server url> <target directory> (default is current directory)")
		fmt.Println("Example: batchocr http://192.168.0.3:8080/ c:\\images")
		return
	}

	// Get all files in the target directory
	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error reading target directory: ", err)
		return
	}

	// Iterate through the files and process each one
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Skipping directory: ", file.Name())
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))

		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			// Process the file
			fmt.Println("Processing file: ", file.Name())
			postFile(filepath.Join(targetDir, file.Name()))
		}
	}
}

type OCRResponse struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
}

func postFile(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", serverUrl, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println("Error posting file: ", err)
		return
	}
	defer res.Body.Close()

	// Read the response
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	ocr := OCRResponse{}
	if err := json.Unmarshal(content, &ocr); err != nil {
		fmt.Println("Error parsing response: ", err)
	}

	if ocr.Success {
		// Write the response to a file
		outFile := strings.Replace(filename, filepath.Ext(filename), ".txt", 1)
		err = ioutil.WriteFile(outFile, []byte(ocr.Text), 0644)
		if err != nil {
			fmt.Println("Error writing response: ", err)
			return
		}
		fmt.Println("Response written to: ", outFile)
	} else {
		fmt.Println("Error processing file: ", ocr.Text)
	}
}
