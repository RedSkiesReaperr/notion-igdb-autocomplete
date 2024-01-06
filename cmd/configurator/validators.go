package main

import (
	"fmt"
	"strconv"
)

func validateString(s string) error {
	if len(s) <= 0 {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}

func validateInteger(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return err
	}
	return nil
}
