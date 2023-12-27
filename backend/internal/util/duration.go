package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Day   = 24 * time.Hour
	Week  = 7 * Day
	Year  = 365 * Day // Simplified, doesn"t account for leap years
	Month = Year / 12 // Simplified, assumes each month is roughly 1/12 of a year
)

// convertToDuration takes a string like "1d" or "1y" and converts it to a time.Duration
func ConvertToDuration(s string) (time.Duration, error) {
	if len(s) < 2 {
		return 0, fmt.Errorf("invalid format")
	}

	numberPart := s[:len(s)-1]
	unit := s[len(s)-1]

	number, err := strconv.Atoi(numberPart)
	if err != nil {
		return 0, err
	}

	switch strings.ToLower(s[len(s)-1:]) {
	case "s":
		return time.Duration(number) * time.Second, nil
	case "m":
		return time.Duration(number) * time.Minute, nil
	case "h":
		return time.Duration(number) * time.Hour, nil
	case "d":
		return time.Duration(number) * 24 * time.Hour, nil
	case "y":
		// Assuming 1 year has 365 days for simplicity
		return time.Duration(number) * 24 * 365 * time.Hour, nil
	default:
		return 0, fmt.Errorf("invalid time unit: %q", unit)
	}
}
