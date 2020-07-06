package main

import (
	"flag"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var telegramUserId int

func main() {
	keyFlag := flag.String("key", "", "telegram api key")
	userIdFlag := flag.Int("user", 0, "telegram user id")

	flag.Parse()

	telegramUserId = *userIdFlag
	bot := getBot(*keyFlag)

	bot.Handle("/add", add)

	bot.Start()
}

func add(m *telebot.Message) {
	if m.Sender.ID != telegramUserId {
		return
	}

	fmt.Printf("%+v/n", m)
}

func getBot(apiKey string) (bot *telebot.Bot) {

	bot, err := telebot.NewBot(telebot.Settings{
		Token: apiKey,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal("Error connecting to telegram, please check your api key.")
	}

	return bot
}
