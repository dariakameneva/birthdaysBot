package main

import (
	"log"
	"strings"

	"gopkg.in/telegram-bot-api.v1"
)

func reply(text string, chatID int, bot *tgbotapi.BotAPI) (string, error) {
	replyMsg := processCommand(text)

	msg := tgbotapi.NewMessage(chatID, replyMsg)
	resMsg, err := bot.Send(msg)
	return resMsg.Text, err
}

func processCommand(text string) string {
	command, ok := commands[getCommandName(text)]
	if !ok {
		return "Unknown command"
	}

	param := getParameter(text)
	resultMessage, err := command.(func(string) (string, error))(param)
	if err != nil {
		log.Println(err.Error())
		return "Something went wrong... Please, try again or contact my creator"
	}
	return resultMessage
}
func getCommandName(fullText string) string {
	return strings.Split(fullText, " ")[0]
}

func getParameter(fullText string) string {
	return strings.TrimPrefix(fullText, getCommandName(fullText)+" ")
}
