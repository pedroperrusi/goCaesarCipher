package main

import "strings"

// replaceAtIndex : https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go
func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func invCesarCypher(nCases uint8, src string) string {
	var dst = strings.ToLower(src)

	for i := 0; i < len(dst); i++ {
		if dst[i] >= 'a' && dst[i] <= 'z' {
			dst = replaceAtIndex(dst, rune(dst[i]-nCases), i)
		}
	}
	return dst
}
