package lineflex

import "github.com/line/line-bot-sdk-go/v7/linebot"

func GenerateStatusRow(name, status string) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Layout:     linebot.FlexBoxLayoutTypeHorizontal,
		PaddingTop: linebot.FlexComponentPaddingTypeLg,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Text:    name,
				Gravity: linebot.FlexComponentGravityTypeCenter,
				Size:    linebot.FlexTextSizeTypeSm,
			},
			&linebot.BoxComponent{
				Layout:          linebot.FlexBoxLayoutTypeVertical,
				BackgroundColor: "#E3F2C1",
				CornerRadius:    linebot.FlexComponentCornerRadiusTypeLg,
				PaddingTop:      linebot.FlexComponentPaddingTypeSm,
				PaddingBottom:   linebot.FlexComponentPaddingTypeSm,
				Width:           "40%",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:  status,
						Align: linebot.FlexComponentAlignTypeCenter,
						Size:  linebot.FlexTextSizeTypeSm,
					},
				},
			},
		},
	}
}

func NewAppIntegrationStatusTemplate() *linebot.BubbleContainer {
	flex0 := 0

	return &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical, // vertical parent
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "Integration Status",
					Size:   linebot.FlexTextSizeTypeXl,
					Weight: linebot.FlexTextWeightTypeBold,
				},
				&linebot.BoxComponent{
					Layout:          linebot.FlexBoxLayoutTypeVertical,
					Contents:        []linebot.FlexComponent{},
					BackgroundColor: "#E8E8E8",
					Height:          "3px",
					Margin:          linebot.FlexComponentMarginTypeMd,
				},
				GenerateStatusRow("123", "123"),
			},
		},
		Footer: &linebot.BoxComponent{
			Layout:  linebot.FlexBoxLayoutTypeVertical,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Style:  linebot.FlexButtonStyleTypeLink,
					Height: linebot.FlexButtonHeightTypeSm,
					Action: &linebot.URIAction{
						Label: "LOGIN",
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
	}
}
