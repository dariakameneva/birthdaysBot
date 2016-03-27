package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v1"
)

var (
	testToken  string
	testChatID int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	testToken = os.Getenv("TEST_TOKEN")

	testChatID, err = strconv.Atoi(os.Getenv("TEST_CHAT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	if testToken == "" {
		log.Fatal("Test token is not defined in .env file")
	}
}

func TestProcessCommand(t *testing.T) {
	testData := map[string]string{
		"/somecmd":   unknownCommand,
		"/set Jan 1": fmt.Sprintf(setBDOkFmtReply, 1, "January"),
		"/start":     passQuestion + passwAddition,
	}

	for msg, expected := range testData {
		result := processCommand(msg)
		if result != expected {
			t.Error(msg + ": Reply message is incorrect: " + result + " instead of " + expected)
		}
	}
}

func TestReply(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(testToken)
	if err != nil {
		t.Error(err)
	}

	testData := []string{
		"/somecmd",
		"/start",
		"/set",
		"/set Jan 1",
		"test text",
	}

	for _, command := range testData {
		if err := reply(command, testChatID, bot); err != nil {
			t.Error(err)
		}
	}
}
