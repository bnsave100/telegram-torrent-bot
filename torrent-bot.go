package main

import (
	"flag"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var telegramUserId int

var bot = func() (bot *telebot.Bot) {
	keyFlag := flag.String("key", "", "telegram api key")
	userIdFlag := flag.Int("user", 0, "telegram user id")

	flag.Parse()

	telegramUserId = *userIdFlag

	bot, err := telebot.NewBot(telebot.Settings{
		Token: *keyFlag,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal("Error connecting to telegram, please check your api key.")
	}

	return bot
}()

func main() {
	bot.Handle("/add", add)

	bot.Start()
}

func add(m *telebot.Message) {
	if m.Sender.ID != telegramUserId {
		return
	}

	fmt.Printf("%+v/n", m)
}
