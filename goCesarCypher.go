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
	description := ProblemCase{}

	if readFromWeb {
		// Read description from HTTPS
		httpsGetJSON(myURL, &description)
	} else {
		// Read description from file
		readJSON(inFileName, &description)
	}

	readPloblemCaseJSON(description)

	writeJSON(outFileName, description)
}
