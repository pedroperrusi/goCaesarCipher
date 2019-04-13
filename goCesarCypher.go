package main

// Read from file hardcoded condition
const readFromWeb = false

// HTTPS address of my JSON file
const myURL = "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

// Input file name
const inFileName = "data/question.json"

// Output file name
const outFileName = "data/answer.json"

func main() {
	// Problem Description
	data := ProblemCase{}

	if readFromWeb {
		// Read description from HTTPS
		httpsGetJSON(myURL, &data)
	} else {
		// Read description from file
		readJSON(inFileName, &data)
	}

	readPloblemCaseJSON(data)

	// Perform Cesar Cypher

	writeJSON(outFileName, data)
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
