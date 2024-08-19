package tui

import "time"

func getProportionOf(number int, proportion float64) int {
	return int(float64(number) * proportion)
}

func timeToHumanDate(t time.Time) string {
	return t.Format("01/02/2006 15:04:05")
}
