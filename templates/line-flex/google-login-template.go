package lineflex

import "github.com/line/line-bot-sdk-go/v7/linebot"

func NewGoogleLoginTemplate(url string) *linebot.BubbleContainer {
	flexValue := 5

	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			URL:         "https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/2008px-Google_%22G%22_Logo.svg.png",
			Size:        "full",
			AspectRatio: "20:13",
			AspectMode:  "fit",
			Action: &linebot.URIAction{
				Label: "WEBSITE",
				URI:   url,
			},
			Align: "center",
		},
		Body: &linebot.BoxComponent{
			Layout: "vertical",
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "Login with Google",
					Weight: "bold",
					Size:   "xl",
				},
				&linebot.BoxComponent{
					Layout:  "vertical",
					Margin:  "lg",
					Spacing: "sm",
					Contents: []linebot.FlexComponent{
						&linebot.BoxComponent{
							Layout:  "baseline",
							Spacing: "sm",
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Text:  "Your line ID will be associated with your google calendar",
									Wrap:  true,
									Color: "#666666",
									Size:  "sm",
									Flex:  &flexValue,
								},
							},
						},
					},
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Layout:  "vertical",
			Spacing: "sm",
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Style:  "link",
					Height: "sm",
					Action: &linebot.URIAction{
						Label: "WEBSITE",
						URI:   url,
					},
				},
				&linebot.BoxComponent{
					Layout:   "vertical",
					Contents: []linebot.FlexComponent{},
					Margin:   "sm",
				},
			},
		},
	}
}
