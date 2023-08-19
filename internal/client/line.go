package client

import (
	"log"
	"os"

	lineflex "github.com/fadelananda/go-line-chatbot/templates/line-flex"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineClient struct {
	lineBotClient *linebot.Client
}

func NewLineClient() (*LineClient, error) {
	lineBotClient, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &LineClient{
		lineBotClient: lineBotClient,
	}, nil
}

func (client *LineClient) GetLineBotInstance() *linebot.Client {
	return client.lineBotClient
}

func (client *LineClient) BroadcastTemplateMessage(templatePath string) {
	templateFileContent, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	flexMessage, err := linebot.UnmarshalFlexMessageJSON(templateFileContent)
	if err != nil {
		log.Fatal(err)
	}

	client.lineBotClient.BroadcastMessage(linebot.NewFlexMessage("new message", flexMessage)).Do()
}

func (client *LineClient) SendTemplateMessage(userId, url string) {
	flex1 := lineflex.NewGoogleLoginTemplate(url)

	client.lineBotClient.PushMessage(userId, linebot.NewFlexMessage("new message", flex1)).Do()
}
