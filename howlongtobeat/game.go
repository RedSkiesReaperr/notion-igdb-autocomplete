package howlongtobeat

import (
	"fmt"
	"math"
	"time"
)

type Result struct {
	Count int   `json:"count,required"`
	Data  Games `json:"data,required"`
}

type Game struct {
	Name           string `json:"game_name,required"`
	CompletionMain int    `json:"comp_main,required"`
	CompletionPlus int    `json:"comp_plus,required"`
	CompletionFull int    `json:"comp_100,required"`
}

type Games []Game

func (g Game) ReadableCompletion(rawCompletion int) string {
	duration := time.Duration(rawCompletion * int(time.Second))
	hours := int(math.Round(duration.Hours()))

	if hours <= 0 {
		return "Not available"
	}

	return fmt.Sprintf("%dh", hours)
}
