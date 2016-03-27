package main

import (
	"strings"

	"gopkg.in/telegram-bot-api.v1"
)

func reply(text string, chatID int, bot *tgbotapi.BotAPI) error {
	replyMsg := processCommand(text)

	msg := tgbotapi.NewMessage(chatID, replyMsg)
	_, err := bot.Send(msg)
	return err
}

func processCommand(text string) string {
	commandName := getCommandName(text)
	param := getParameter(text)

	command, ok := commands[commandName]
	if !ok {
		return unknownCommand
	}

	return command.(func(string) string)(param)
}
func getCommandName(fullText string) string {
	if !strings.HasPrefix(fullText, "/") {
		return ""
	}
	return strings.Split(fullText, " ")[0]
}

func getParameter(fullText string) string {
	res := strings.TrimPrefix(fullText, getCommandName(fullText))
	return strings.TrimSpace(res)
}
