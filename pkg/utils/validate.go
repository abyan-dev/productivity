package utils

import (
	"regexp"
	"strings"
	"time"
)

func ValidateEmail(email string) (bool, string) {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	if !re.MatchString(email) {
		if !strings.Contains(email, "@") {
			return false, "Email must contain an '@' symbol"
		}
		if !strings.Contains(email, ".") {
			return false, "Email must contain a '.' symbol"
		}
		if strings.Contains(email, " ") {
			return false, "Email must not contain spaces"
		}
		return false, "Email format is invalid"
	}
	return true, "Email is valid"
}

func ValidateTime(timeStr string) (bool, string, time.Time) {
	const timeLayout = "2006-01-02T15:04:05Z07:00"

	parsedTime, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		return false, "Time format is invalid", time.Time{}
	}

	return true, "Time is valid", parsedTime
}
