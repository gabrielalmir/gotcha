package utils

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidURL = errors.New("invalid URL format")
)

func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(/.*)?$`)
	return re.MatchString(url)
}
