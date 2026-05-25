package howlongtobeat

import (
	"fmt"
	"math"
	"time"

	"github.com/RedSkiesReaperr/howlongtobeat"
	"github.com/jomei/notionapi"
)

type GameWrapper struct {
	Game *howlongtobeat.Game
}

func (ug GameWrapper) readableCompletion(rawCompletion int) string {
	duration := time.Duration(rawCompletion * int(time.Second))
	hours := int(math.Round(duration.Hours()))

	if hours <= 0 {
		return "Not available"
	}

	return fmt.Sprintf("%dh", hours)
}

// Implements notion.PageUpdateRequester interface
func (ug GameWrapper) UpdateRequest() notionapi.PageUpdateRequest {
	return notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Time to complete (Main Story)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: ug.readableCompletion(ug.Game.CompletionMain),
						},
					},
				},
			},
			"Time to complete (Main + Sides)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: ug.readableCompletion(ug.Game.CompletionPlus),
						},
					},
				},
			},
			"Time to complete (Completionist)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: ug.readableCompletion(ug.Game.CompletionFull),
						},
					},
				},
			},
		},
	}
}
