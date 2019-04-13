package main

import "testing"

func TestReplace(t *testing.T) {
	str := "aaaaaaa"
	dst := replaceAtIndex(str, 'b', 3)
	expected := "aaabaaa"
	if dst != expected {
		t.Errorf("Replace at Index was incorrect, got: %s, want: %s.", dst, expected)
	}
}

func TestCesarCypher(t *testing.T) {
	str := "cccc"
	var cases uint8 = 2
	dst := invCesarCypher(cases, str)
	expected := "aaaa"
	if dst != expected {
		t.Errorf("Invert Cesar Cypher is incorrect, got: %s, want: %s.", dst, expected)
	}
}
