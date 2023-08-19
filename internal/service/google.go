package service

import (
	"errors"

	"github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
)

type GoogleService struct {
	client *client.GoogleCalendarClient
}

func NewGoogleService(client *client.GoogleCalendarClient) *GoogleService {
	return &GoogleService{
		client: client,
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

	token, err := s.client.ExchangeOauthCode(code)
	if err != nil {
		return errors.New("cannot exchange code for token")
	}

	s.client.ListEvent(token)
	return nil
}
