package main

func main() {
	// HTTPS address of my JSON file
	var myURL = "https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=b566cecbba033c449a449a8d3e16fb69e10bddcc"

	// Read problem case from the url file
	description := ProblemCase{}
	getJSON(myURL, &description)

	readPloblemCaseJSON(description)

	writeJSON("data/answer.json", description)
}
