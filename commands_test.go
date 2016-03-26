package main

import "testing"

func TestSetBirthday(t *testing.T) {
	res, err := setBirthday("Jan 1")
	if err != nil {
		t.Error("Error when setting a birthday:", err.Error())
	}
	if res != "Saved your birthday: 1 of January" {
		t.Error("Incorrect reply string:", res)
	}
}
