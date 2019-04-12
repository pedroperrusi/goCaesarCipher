package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HTTPS client timout definition
var myClient = &http.Client{Timeout: 10 * time.Second}

// Simple HTTPS client to parse JSON
func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
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
