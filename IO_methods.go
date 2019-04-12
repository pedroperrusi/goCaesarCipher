package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// ProblemCase : JSON description of the problem case to be solved.
type ProblemCase struct {
	NumeroCasas         int    `json:"numero_casas"`
	Token               string `json:"token"`
	Cifrado             string `json:"cifrado"`
	Decifrado           string `json:"decifrado"`
	ResumoCriptografado string `json:"resumo_criptografado"`
}

// readPloblemCaseJSON : prints all data types from ProblemCase structure
func readPloblemCaseJSON(problemJSON ProblemCase) {
	// Read all JSON
	print("NumeroCasas: ")
	println(problemJSON.NumeroCasas)

	print("Token: ")
	println(problemJSON.Token)

	print("Cifrado: ")
	println(problemJSON.Cifrado)

	print("Decifrado: ")
	println(problemJSON.Decifrado)

	print("ResumoCriptografado: ")
	println(problemJSON.ResumoCriptografado)
}
