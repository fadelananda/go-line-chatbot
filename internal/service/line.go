package service

import (
	"fmt"

	clients "github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineService struct {
	LineClient           *clients.LineClient
	GoogleCalendarClient *clients.GoogleCalendarClient
}

func NewLineService(lineClient *clients.LineClient, googleCalendarClient *clients.GoogleCalendarClient) *LineService {
	return &LineService{
		LineClient:           lineClient,
		GoogleCalendarClient: googleCalendarClient,
	}
}

func (s *LineService) HandleWebhookEvents(events []*linebot.Event) {
	utils.LogInfo("handle webhook events", map[string]interface{}{
		"events": events,
	})

	// TODO: get google url
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message.Text)
				userId := event.Source.UserID
				oauthURL, _ := s.GoogleCalendarClient.GenerateOauthURL(userId)

				if message.Text == "login" {
					s.LineClient.SendTemplateMessage(userId, oauthURL)
				}
			}
		}
	}
}
