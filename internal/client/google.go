package client

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarClient struct {
	config *oauth2.Config
}

func NewGoogleCalendarClient() (*GoogleCalendarClient, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	return &GoogleCalendarClient{
		config: config,
	}, nil
}

func (client *GoogleCalendarClient) GenerateOauthURL(lineId string) (string, error) {
	jwtValue, err := utils.GenerateOauthJWT(lineId)
	if err != nil {
		return "", err
	}

	return client.config.AuthCodeURL(jwtValue, oauth2.AccessTypeOffline), nil
}

func (client *GoogleCalendarClient) ExchangeOauthCode(code string) (*oauth2.Token, error) {
	tok, err := client.config.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func (client *GoogleCalendarClient) ListEvent(tok *oauth2.Token) (*calendar.Events, error) {
	ctx := context.Background()
	clientConf := client.config.Client(ctx, tok)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(clientConf))
	if err != nil {
		utils.LogError("unable to retrieve calendar client", err, map[string]interface{}{
			"token": tok,
		})
		return nil, err
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	return events, nil
}
