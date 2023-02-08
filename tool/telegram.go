package tool

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

var config Config
var bot *tgbotapi.BotAPI

func init() {
	config.Init()
}

type Config struct {
	Token  string `yaml:"token"`
	ChatId int64  `yaml:"chat-id"`
	Url    string `yaml:"url"`
}

func (c *Config) Init() {
	yamlFile, err := ioutil.ReadFile("telegram-config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		return
	}
	if c.Url != "" {
		bot, err = tgbotapi.NewBotAPIWithClient(c.Token, c.Url, &http.Client{})
	} else {
		bot, err = tgbotapi.NewBotAPI(c.Token)
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func SendMsgText(text string) {
	if bot == nil {
		log.Println("bot is nil")
		return
	}
	message := tgbotapi.NewMessage(config.ChatId, text)
	_, err := bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
