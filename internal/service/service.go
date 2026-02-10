package service

import (
	"Yp_Sprint6/pkg/morse"
	"errors"
	"strings"
)

func AutoDetect(input string) (string, error) {

	if len(input) == 0 {
		return "", errors.New("Empty stroke")
	}

	if isMorseCode(input) {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}

func isMorseCode(s string) bool {

	for _, char := range s {

		if !(char == '.' || char == '-' ||
			char == ' ' || char == '/' ||
			char == '\n' || char == '\r' || char == '\t') {
			return false
		}
	}

	if !strings.ContainsAny(s, ".-") {
		return false
	}

	return true
}
