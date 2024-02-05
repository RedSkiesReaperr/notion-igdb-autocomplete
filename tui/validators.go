package tui

import (
	"fmt"
	"strconv"
)

func ValidateString(s string) error {
	if len(s) <= 0 {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}

func ValidateInteger(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return err
	}
	return nil
}
