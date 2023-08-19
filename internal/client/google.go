package client

import (
	"context"
	"fmt"
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

func (client *GoogleCalendarClient) ListEvent(tok *oauth2.Token) {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client.config.Client(context.Background(), tok)))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}
