package service

import (
	"fmt"
	"time"

	"github.com/fadelananda/go-line-chatbot/entity"
	clients "github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	lineflex "github.com/fadelananda/go-line-chatbot/templates/line-flex"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/oauth2"
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
	utils.LogInfo("handle webhook events", map[string]interface{}{
		"events": events,
	})

	// TODO: get google url
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userId := event.Source.UserID

				// parse message
				switch message.Text {
				case "login":
					oauthURL, _ := s.GoogleCalendarClient.GenerateOauthURL(userId)
					loginTemplate := lineflex.NewGoogleLoginTemplate(oauthURL)
					s.LineClient.SendTemplateMessage(userId, "login url", loginTemplate)

				case "list":
					user, err := s.AWSClient.GetDataByLineId(userId)
					if err != nil {
					}

					events, err := s.GoogleCalendarClient.ListEvent(user.AuthToken)
					if err != nil {
					}

					calendarTemplate := lineflex.NewGoogleCalendarList("tes13", events)
					fmt.Println("::::::")
					fmt.Println(userId)
					s.LineClient.SendTemplateMessage(userId, "calendar list", calendarTemplate)

				case "status":
					updateUser := entity.User{
						LineId: userId,
						Email:  "fadelananda124563@gmail.com",
						AuthToken: &oauth2.Token{
							Expiry:       time.Now().Add(time.Hour * 456),
							AccessToken:  "jahahaha",
							RefreshToken: "1//0gcgFBVOyPeSHCgYIARAAGBASNwF-L9IrcRKvkADNKk4jLrSKrQ1Niee3a94VcSEFIWHXfDlaAcjvg4cNNQJ9ZE8Wk8ZOwV-HV40",
							TokenType:    "Bearer",
						},
					}
					s.AWSClient.UpdateUser(userId, updateUser)
				}
			}
		}
	}
}
