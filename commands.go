package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var (
	commands = map[string]interface{}{
		"/start": askPass,
		"/passw": askPass,
		"/set":   setBirthday,
		"/help":  sendHelp,
	}
)

func askPass(param string) string {
	if param == "" {
		return passQuestion + passwAddition
	}
	if checkPass(param) {
		return passwOkReply
	}
	return passwFailReply
}

func sendHelp(param string) string {
	return helpReply
}

func checkPass(param string) bool {
	possibleAnswers := strings.Split(passAnswer, ";")
	for _, ans := range possibleAnswers {
		if strings.Contains(strings.ToLower(param), strings.ToLower(ans)) {
			return true
		}
	}
	return false
}

func setBirthday(param string) string {
	date, err := time.Parse("Jan 2", param)
	if err != nil {
		if strings.Contains(err.Error(), "day out of range") {
			return setBDOutOfRangeReply
		}
		log.Println(err)
		return setBDFailReply
	}
	return fmt.Sprintf(setBDOkFmtReply, date.Day(), date.Month().String())
}
