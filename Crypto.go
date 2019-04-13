package main

import (
	"crypto/sha1"
	"strings"
)

// replaceAtIndex : https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go
func replaceAtIndex(str string, replacement uint8, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

// Unencript a single char accordingly to Cesar Cypher
func unencryptChar(nCases uint8, char uint8) uint8 {
	var newChar = uint8(char - nCases)
	if newChar < 'a' {
		newChar += 'z' - 'a' + 1
	}
	return newChar
}

// Performs Cesar Cypher to a String
func invCesarCypher(nCases uint8, src string) string {
	var dst = strings.ToLower(src)

	for i := 0; i < len(dst); i++ {
		if dst[i] >= 'a' && dst[i] <= 'z' {
			var newChar = unencryptChar(nCases, dst[i])
			dst = replaceAtIndex(dst, newChar, i)
		}
	}
	return dst
}

// Creates a SHA1 hash of a string
func stringToSHA1(src string) []byte {
	h := sha1.New()
	h.Write([]byte(src))
	return h.Sum(nil)
}
