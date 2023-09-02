package service

import (
	"time"

	clients "github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	lineflex "github.com/fadelananda/go-line-chatbot/templates/line-flex"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineService struct {
	LineClient           *clients.LineClient
	GoogleCalendarClient *clients.GoogleCalendarClient
	AWSClient            *clients.AWSClient
}

func NewLineService(lineClient *clients.LineClient, googleCalendarClient *clients.GoogleCalendarClient, awsClient *clients.AWSClient) *LineService {
	return &LineService{
		LineClient:           lineClient,
		GoogleCalendarClient: googleCalendarClient,
		AWSClient:            awsClient,
	}
}

func (s *LineService) HandleWebhookEvents(events []*linebot.Event) {
	utils.LogInfo("handling webhook events", map[string]interface{}{
		"events": events,
	})

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userId := event.Source.UserID

				// parse message
				// TODO: modularize all parsing message
				switch message.Text {
				case "help":
					s.LineClient.SendTextMessage(userId, "list of commands available: \n- login \n- list \n- help")

				case "login":
					oauthURL, _ := s.GoogleCalendarClient.GenerateOauthURL(userId)
					loginTemplate := lineflex.NewGoogleLoginTemplate(oauthURL)
					s.LineClient.SendTemplateMessage(userId, "login url", loginTemplate)

				case "list":
					user, err := s.AWSClient.GetDataByLineId(userId)
					if err != nil {
					}

					events, err := s.GoogleCalendarClient.ListEvent(user.AuthToken, time.Time{})
					if err != nil {
					}

					calendarTemplate := lineflex.NewGoogleCalendarList("tes13", events)
					s.LineClient.SendTemplateMessage(userId, "calendar list", calendarTemplate)

				// TODO: DETERMINE IF NEED TO HANDLE REFRESH BY OURSELF

				case "status":
					user, _ := s.AWSClient.GetDataByLineId("userId123123132")
					if user.IsEmpty() {
						oauthURL, _ := s.GoogleCalendarClient.GenerateOauthURL(userId)
						loginTemplate := lineflex.NewGoogleLoginTemplate(oauthURL)
						s.LineClient.SendTextMessage(userId, "You are not logged in yet, please login using the link below!")
						s.LineClient.SendTemplateMessage(userId, "login url", loginTemplate)
						break
					}
					template := lineflex.NewAppIntegrationStatusTemplate()
					s.LineClient.SendTemplateMessage(userId, "status", template)
				}
			}
		}
	}
}
