package main

import (
	"encoding/hex"
	"fmt"
)

// Read from file hardcoded condition
const readFromWeb = true

// Post to web hardcoded condition
const postToWeb = true

// HTTPS address of question.json to be read
const sourceURL = "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

// HTTPS target address for the answer.json to be sent
const targetURL = "https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

// Input file name
const inFileName = "data/question.json"

// Output file name
const outFileName = "data/answer.json"

func main() {
	// Problem Description
	data := ProblemCase{}

	if readFromWeb {
		// Read description from HTTPS
		httpsGetJSON(sourceURL, &data)
	} else {
		// Read description from file
		readJSON(inFileName, &data)
	}

	println("\nProblem case read:\n")
	readPloblemCaseJSON(data)

	// Perform Cesar Cypher
	data.Decifrado = invCesarCypher(data.NumeroCasas, data.Cifrado)

	// Generate hash resume
	data.ResumoCriptografado = hex.EncodeToString(stringToSHA1(data.Decifrado))

	println("\nProblem case solved:\n")
	readPloblemCaseJSON(data)

	writeJSON(outFileName, data)

	// Send answer.json by POST
	if postToWeb {
		postFile(outFileName, targetURL)
	}
}

// ProblemCase : JSON description of the problem case to be solved.
type ProblemCase struct {
	NumeroCasas         uint8  `json:"numero_casas"`
	Token               string `json:"token"`
	Cifrado             string `json:"cifrado"`
	Decifrado           string `json:"decifrado"`
	ResumoCriptografado string `json:"resumo_criptografico"`
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
	fmt.Printf("%s\n", problemJSON.ResumoCriptografado)
}
