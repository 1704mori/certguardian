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
	Year  = 365 * Day // Simplified, doesn't account for leap years
	Month = Year / 12 // Simplified, assumes each month is roughly 1/12 of a year
)

// convertToDuration takes a string like "1d" or "1y" and converts it to a time.Duration
func ConvertToDuration(s string) (time.Duration, error) {
	if len(s) < 2 {
		return 0, fmt.Errorf("invalid format")
	}

	num, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, err
	}

	switch strings.ToLower(s[len(s)-1:]) {
	case "d":
		return time.Duration(num) * Day, nil
	case "y":
		return time.Duration(num) * Year, nil
	case "s":
		return time.Duration(num) * time.Second, nil
	case "h":
		return time.Duration(num) * time.Hour, nil
	case "w":
		return time.Duration(num) * Week, nil
	case "m":
		return time.Duration(num) * Month, nil
	default:
		return 0, fmt.Errorf("unknown time unit")
	}
}
