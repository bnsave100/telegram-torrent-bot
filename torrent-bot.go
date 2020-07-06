package main

import (
	"flag"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
	"time"
	"torrent-bot/qbittorrent"
)

var telegramUserId int
var qBittorrentClient *qbittorrent.QBittorrent

func main() {
	keyFlag := flag.String("key", "", "telegram api key")
	userIdFlag := flag.Int("user", 0, "telegram user id")
	qblFlag := flag.String("qbl", "", "qBittorrent username")
	qbpFlag := flag.String("qbp", "", "qBittorrent password")
	qbuFlag := flag.String("qbu", "", "qBittorrent base url")

	flag.Parse()

	telegramUserId = *userIdFlag
	qBittorrentClient = qbittorrent.NewQBittorrent(*qblFlag, *qbpFlag, *qbuFlag)

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
