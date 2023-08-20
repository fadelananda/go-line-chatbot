package service

import (
	"errors"

	"github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
)

type GoogleService struct {
	gClient   *client.GoogleCalendarClient
	awsClient *client.AWSClient
}

func NewGoogleService(gClient *client.GoogleCalendarClient, awsClient *client.AWSClient) *GoogleService {
	return &GoogleService{
		gClient:   gClient,
		awsClient: awsClient,
	}
}

func (s *GoogleService) HandleOauthCallback(code, state string) error {
	claims, err := utils.ValidateOauthJWT(state)
	if err != nil {
		return errors.New("cannot validate state")
	}

	_, ok := claims["line_id"].(string)
	if !ok {
		return errors.New("error when extracting line_id from claim")
	}

	token, err := s.gClient.ExchangeOauthCode(code)
	if err != nil {
		return errors.New("cannot exchange code for token")
	}

	s.awsClient.AddUser(client.User{
		LineId:    claims["line_id"].(string),
		AuthToken: token,
	})
	s.gClient.ListEvent(token)
	return nil
}
