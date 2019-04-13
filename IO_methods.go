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
	"time"
)

// HTTPS client timout definition
var myClient = &http.Client{Timeout: 10 * time.Second}

// Simple HTTPS client to parse JSON
func httpsGetJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// readJSON : Loads an structure with JSON from a file
func readJSON(filename string, target interface{}) {
	fileData, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(fileData), &target)
}

// Write a structure type to a JSON file
func writeJSON(filename string, target interface{}) {
	targetJSON, _ := json.Marshal(target)
	ioutil.WriteFile(filename, targetJSON, 0644)
}

// Upload file by POST to a URL (thanks https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html)
func postFile(filename string, targetURL string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("answer", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}
