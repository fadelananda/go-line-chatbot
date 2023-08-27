package lineflex

import (
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"google.golang.org/api/calendar/v3"
)

func generateRowData(eventTime, eventName, meetingLink string) *linebot.BoxComponent {
	flex1 := 1
	flex4 := 4

	return &linebot.BoxComponent{
		Layout: linebot.FlexBoxLayoutTypeVertical, // vertical per row need to separate
		Margin: linebot.FlexComponentMarginTypeLg,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Layout: linebot.FlexBoxLayoutTypeHorizontal, // horizontal top
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Layout:          linebot.FlexBoxLayoutTypeHorizontal,
						Contents:        []linebot.FlexComponent{},
						BackgroundColor: "#E8E8E8",
						Width:           "5px",
						Margin:          linebot.FlexComponentMarginTypeNone,
						Flex:            &flex1,
					},
					&linebot.TextComponent{
						Text:        eventTime,
						Color:       "#BBBFCA",
						Flex:        &flex1,
						OffsetStart: linebot.FlexComponentOffsetTypeMd,
						Weight:      linebot.FlexTextWeightTypeBold,
					},
					&linebot.TextComponent{
						Text:        eventName,
						Wrap:        false,
						Flex:        &flex4,
						OffsetStart: linebot.FlexComponentOffsetTypeXl,
					},
				},
			},
			&linebot.BoxComponent{
				Layout: linebot.FlexBoxLayoutTypeHorizontal, // horizontal bottom
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Layout:          linebot.FlexBoxLayoutTypeHorizontal,
						Contents:        []linebot.FlexComponent{},
						BackgroundColor: "#E8E8E8",
						Width:           "5px",
						Margin:          linebot.FlexComponentMarginTypeNone,
						Flex:            &flex1,
					},
					&linebot.TextComponent{
						Text:        "meet",
						Color:       "#BBBFCA",
						Size:        linebot.FlexTextSizeTypeSm,
						Flex:        &flex1,
						OffsetStart: linebot.FlexComponentOffsetTypeMd,
						Weight:      linebot.FlexTextWeightTypeRegular,
					},
					&linebot.TextComponent{
						Text:        meetingLink,
						Wrap:        false,
						Size:        linebot.FlexTextSizeTypeSm,
						Flex:        &flex4,
						OffsetStart: linebot.FlexComponentOffsetTypeXl,
					},
				},
			},
		},
	}
}

func NewGoogleCalendarList(date string, events *calendar.Events) *linebot.BubbleContainer {
	flex0 := 0

	bodyContent := []linebot.FlexComponent{}
	for _, event := range events.Items {
		dateTime := event.Start.DateTime

		if dateTime == "" {
			dateTime = "All Day"
		} else {
			parsedDateTime, err := time.Parse(time.RFC3339, dateTime)
			if err != nil {
				fmt.Println("Error parsing datetime:", err)
			}
			timeOnly := parsedDateTime.Format("15:04")
			dateTime = timeOnly
		}
		row := generateRowData(dateTime, event.Summary, "https://meet.google.com")
		bodyContent = append(bodyContent, row)
	}

	return &linebot.BubbleContainer{
		Header: &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "Event List",
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXl,
				},
				&linebot.TextComponent{
					Text:   date,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeMd,
				},
			},
		},
		Body: &linebot.BoxComponent{
			Layout:   linebot.FlexBoxLayoutTypeVertical, // vertical parent
			Contents: bodyContent,
		},
		Footer: &linebot.BoxComponent{
			Layout:  linebot.FlexBoxLayoutTypeVertical,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Style:  linebot.FlexButtonStyleTypeLink,
					Height: linebot.FlexButtonHeightTypeSm,
					Action: &linebot.URIAction{
						Label: "Open in Google Calendar",
						URI:   "https://google.com",
					},
				},
				&linebot.BoxComponent{
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{},
					Margin:   linebot.FlexComponentMarginTypeSm,
				},
			},
			Flex: &flex0,
		},
		Styles: &linebot.BubbleStyle{
			Header: &linebot.BlockStyle{
				BackgroundColor: "#F4F4F2",
			},
		},
	}
}
