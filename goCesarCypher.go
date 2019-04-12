package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// ProblemCase : JSON description of the problem case to be solved.
type ProblemCase struct {
	NumeroCasas         int
	Token               string
	Cifrado             string
	Decifrado           string
	ResumoCriptografado string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	// HTTPS address of my JSON file
	var myURL = "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

	problemJSON := ProblemCase{}
	getJSON(myURL, &problemJSON)

	// Read all JSON
	print("NumeroCasas: ")
	println(problemJSON.NumeroCasas + 1)

	print("Token: ")
	println(problemJSON.Token)

	print("Cifrado: ")
	println(problemJSON.Cifrado)

	print("Decifrado: ")
	println(problemJSON.Decifrado)

	print("ResumoCriptografado: ")
	println(problemJSON.ResumoCriptografado)
}
