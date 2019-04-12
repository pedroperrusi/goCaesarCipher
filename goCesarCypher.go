package main

func main() {
	// HTTPS address of my JSON file
	var myURL = "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

	// Read problem case from the url file
	problemJSON := ProblemCase{}
	getJSON(myURL, &problemJSON)

	// numero_casas is harcoded as its right value :(
	problemJSON.NumeroCasas = 1

	readPloblemCaseJSON(problemJSON)
}
