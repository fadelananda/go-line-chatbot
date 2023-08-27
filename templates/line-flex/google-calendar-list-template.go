package lineflex

import (
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"google.golang.org/api/calendar/v3"
)

// TODO: add feature
// google meet link

func generateRowData(eventTime, eventName, meetingLink, separatorColor string) *linebot.BoxComponent {
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
						BackgroundColor: separatorColor,
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
						BackgroundColor: separatorColor,
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
	flex1 := 1
	flex5 := 5

	defaultEventSeparatorColor := "#E8E8E8"
	allDayEventSeparatorColor := "#EEE0C9"

	var separatorColor string

	bodyContent := []linebot.FlexComponent{}
	for _, event := range events.Items {
		separatorColor = defaultEventSeparatorColor
		dateTime := event.Start.DateTime

		if dateTime == "" {
			separatorColor = allDayEventSeparatorColor
			dateTime = "*"
		} else {
			parsedDateTime, err := time.Parse(time.RFC3339, dateTime)
			if err != nil {
				fmt.Println("Error parsing datetime:", err)
			}
			timeOnly := parsedDateTime.Format("15:04")
			dateTime = timeOnly
		}
		row := generateRowData(dateTime, event.Summary, "https://meet.google.com", separatorColor)
		bodyContent = append(bodyContent, row)
	}

	bodyContent = append(bodyContent, &linebot.BoxComponent{
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Margin: linebot.FlexComponentMarginTypeXxl,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Layout:         linebot.FlexBoxLayoutTypeVertical,
						Flex:           &flex1,
						JustifyContent: linebot.FlexComponentJustifyContentTypeCenter,
						Contents: []linebot.FlexComponent{
							&linebot.BoxComponent{
								Layout:          linebot.FlexBoxLayoutTypeVertical,
								Contents:        []linebot.FlexComponent{},
								BackgroundColor: allDayEventSeparatorColor,
								Width:           "15px",
								Height:          "15px",
								Flex:            &flex1,
								CornerRadius:    linebot.FlexComponentCornerRadiusTypeSm,
							},
						},
					},
					&linebot.TextComponent{
						Text:      "All Day",
						Flex:      &flex5,
						Size:      linebot.FlexTextSizeTypeXs,
						Weight:    linebot.FlexTextWeightTypeBold,
						OffsetEnd: linebot.FlexComponentOffsetTypeXxl,
					},
				},
			},
		},
	})

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
