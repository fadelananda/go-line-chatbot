package service

import (
	"errors"
	"time"

	"github.com/fadelananda/go-line-chatbot/entity"
	"github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/repository"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
)

type GoogleService struct {
	gClient    *client.GoogleCalendarClient
	userStorer repository.UserStorer
}

func NewGoogleService(gClient *client.GoogleCalendarClient, userStorer repository.UserStorer) *GoogleService {
	return &GoogleService{
		gClient:    gClient,
		userStorer: userStorer,
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

	s.userStorer.AddUser(entity.User{
		LineId:    claims["line_id"].(string),
		AuthToken: token,
	})
	return nil
}

func (s *GoogleService) ValidateOauthSession(lineId string) {
	user, err := s.userStorer.GetUserById(lineId)
	if err != nil {
		// TODO: error handling
	}

	if user.AuthToken.Expiry.Before(time.Now()) {

	}
}

func (s *GoogleService) RefreshGoogleOauthToken(user entity.User) {
	newToken, err := s.gClient.RefreshOauthToken(user.AuthToken)
	if err != nil {
		// TODO: error handling
	}

	updateUserData := entity.User{
		LineId:    user.LineId,
		AuthToken: newToken,
	}
	s.userStorer.AddUser(updateUserData)
}
