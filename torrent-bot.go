package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"strings"
	"time"
	"torrent-bot/qbittorrent"
)

var telegramApiKey string
var telegramUserId int
var qBittorrentClient *qbittorrent.QBittorrent
var bot *telebot.Bot

func main() {
	keyFlag := flag.String("key", "", "telegram api key")
	userIdFlag := flag.Int("user", 0, "telegram user id")
	qblFlag := flag.String("qbl", "", "qBittorrent username")
	qbpFlag := flag.String("qbp", "", "qBittorrent password")
	qbuFlag := flag.String("qbu", "http://localhost:8080", "qBittorrent base url")

	flag.Parse()

	telegramApiKey = *keyFlag
	telegramUserId = *userIdFlag
	qBittorrentClient = qbittorrent.NewQBittorrent(*qblFlag, *qbpFlag, *qbuFlag)

	bot = getBot(telegramApiKey)

	bot.Handle("/add", add)

	bot.Handle("/list", list)

	bot.Handle(telebot.OnDocument, onFile)

	bot.Start()
}

func add(m *telebot.Message) {
	if m.Sender.ID != telegramUserId && telegramUserId != 0 {
		return
	}

	message := strings.Replace(m.Text, "/add ", "", -1)
	err := qBittorrentClient.Add(strings.Split(message, "\n"))

	if err == nil {
		_, _ = bot.Send(m.Sender, "Success!")
		return
	}

	_, _ = bot.Send(m.Sender, err.Error())
}

func list(m *telebot.Message) {
	torrentList, err := qBittorrentClient.List()

	if err != nil {
		_, _ = bot.Send(m.Sender, err.Error())
		return
	}

	_, _ = bot.Send(m.Sender, torrentList.ToString())
}

func onFile(m *telebot.Message) {
	if m.Sender.ID != telegramUserId && telegramUserId != 0 {
		return
	}

	if m.Document.MIME != "application/x-bittorrent" {
		return
	}

	file := m.Document.MediaFile()
	url := getFileUrl(file)
	err := qBittorrentClient.Add([]string{url})

	if err == nil {
		_, _ = bot.Send(m.Sender, "Success!")
		return
	}

	_, _ = bot.Send(m.Sender, err.Error())
}

func getFileUrl(file *telebot.File) string {
	fileJsonUrl := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", telegramApiKey, file.FileID)

	client := &http.Client{}
	r, _ := client.Get(fileJsonUrl)

	var jsonParsed struct {
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}

	_ = json.NewDecoder(r.Body).Decode(&jsonParsed)
	_ = r.Body.Close()

	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", telegramApiKey, jsonParsed.Result.FilePath)
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
