package main

import (
	"fmt"
	"time"
)

var (
	commands = map[string]interface{}{
		"/set": setBirthday,
	}
)

func setBirthday(param string) (string, error) {
	date, err := time.Parse("Jan 2", param)
	return fmt.Sprintf("Saved your birthday: %d of %s", date.Day(), date.Month().String()), err
}
