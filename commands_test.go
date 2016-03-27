package main

import (
	"fmt"
	"testing"
)

func TestSetBirthday(t *testing.T) {
	testData := map[string]string{
		"Jan 1":  fmt.Sprintf(setBDOkFmtReply, 1, "January"),
		"jan 31": fmt.Sprintf(setBDOkFmtReply, 31, "January"),
		"Feb 31": setBDOutOfRangeReply,
		"31 jan": setBDFailReply,
		"31-jan": setBDFailReply,
		"Jab 31": setBDFailReply,
	}

	for input, expected := range testData {
		res := setBirthday(input)
		if res != expected {
			t.Error(input+":", "Incorrect reply string:", res)
		}
	}
}

func TestCheckPass(t *testing.T) {
	testData := map[string]bool{
		"yes":  true,
		"YES":  true,
		"Yes":  true,
		"はい":   true,
		"Oui!": true,
		"Jaaa": true,
		"Да.":  true,
		"I do certainly think so, yes": true,
		"no!": false,
		"I don't consider it to be true, no": false,
	}

	for input, expected := range testData {
		res := checkPass(input)
		if res != expected {
			t.Error("Unexpected return value on", input)
		}
	}
}

func TestAskPass(t *testing.T) {
	testData := map[string]string{
		"yes": passwOkReply,
		"no":  passwFailReply,
	}

	for input, expected := range testData {
		res := askPass(input)
		if res != expected {
			t.Error("Unexpected return value on", input)
		}
	}
}
