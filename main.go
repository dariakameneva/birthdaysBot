package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/telegram-bot-api.v1"
)

var (
	token        = ""
	debug        = true
	passQuestion = ""
	passAnswer   = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token = os.Getenv("BD_BOT_TOKEN")

	debugParam := strings.ToLower(os.Getenv("DEBUG"))
	if debugParam != "true" && debugParam != "false" {
		log.Fatal("Invalid DEBUG parameter value: must be TRUE or FALSE")
	}
	debug = (debugParam == "TRUE")

	passQuestion = os.Getenv("PASSQUESTION")
	passAnswer = os.Getenv("PASSANSWER")
}

func main() {
	logfile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(logfile)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if err := reply(update.Message.Text, update.Message.Chat.ID, bot); err != nil {
			log.Println(err)
		}
	}
}
