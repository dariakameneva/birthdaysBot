package main

import "testing"

const (
	testToken  = "102140718:AAHkS984gBQG0N7RNC1OVycWwGZVqz-qL-w"
	testChatID = 87956259
)

func TestProcessCommand(t *testing.T) {
	testData := map[string]string{
		"/set Jan 1": "Saved your birthday: 1 of January",
	}

	for msg, replyMsg := range testData {
		result := processCommand(msg)
		if result != replyMsg {
			t.Error("Reply message is incorrect")
		}
	}
}
