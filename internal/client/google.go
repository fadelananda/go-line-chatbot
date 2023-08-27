package client

import (
	"context"
	"fmt"
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

type googleClientError struct {
	FunctionName string
	Err          error
}

func (e *googleClientError) Error() string {
	return fmt.Sprintf("Google client error from %s, error: %v", e.FunctionName, e.Err)
}

func newGoogleClientError(functionName string, err error) *googleClientError {
	return &googleClientError{
		FunctionName: functionName,
		Err:          err,
	}
}

func NewGoogleCalendarClient() (*GoogleCalendarClient, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, newGoogleClientError("NewGoogleCalendarClient", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, newGoogleClientError("NewGoogleCalendarClient", err)
	}

	return &GoogleCalendarClient{
		config: config,
	}, nil
}

func (client *GoogleCalendarClient) GenerateOauthURL(lineId string) (string, error) {
	jwtValue, err := utils.GenerateOauthJWT(lineId)
	if err != nil {
		return "", newGoogleClientError("GenerateOauthURL", err)
	}

	return client.config.AuthCodeURL(jwtValue, oauth2.AccessTypeOffline), nil
}

func (client *GoogleCalendarClient) ExchangeOauthCode(code string) (*oauth2.Token, error) {
	tok, err := client.config.Exchange(context.TODO(), code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, newGoogleClientError("ExchangeOauthCode", err)
	}

	return tok, nil
}

func (client *GoogleCalendarClient) RefreshOauthToken(oldToken *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := client.config.TokenSource(context.Background(), oldToken).Token()
	if err != nil {
		return nil, newGoogleClientError("RefreshOauthToken", err)
	}

	return newToken, nil
}

func (client *GoogleCalendarClient) ListEvent(tok *oauth2.Token, endTime time.Time) (*calendar.Events, error) {
	ctx := context.Background()
	clientConf := client.config.Client(ctx, tok)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(clientConf))
	if err != nil {
		utils.LogError("unable to retrieve calendar client", err, map[string]interface{}{
			"token": tok,
		})
		return nil, newGoogleClientError("ListEvent", err)
	}

	if !endTime.IsZero() {
		utils.LogInfo("zewooo", nil)
	}
	t := time.Now().Format(time.RFC3339)
	startOfDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	endOfDay := startOfDay.Add(24 * time.Hour).Format(time.RFC3339)

	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).TimeMax(endOfDay).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		newGoogleClientError("ListEvent", err)
	}

	return events, nil
}
