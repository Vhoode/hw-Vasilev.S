package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	errMsg := ErrInvalidString

	if len(input) > 0 && unicode.IsDigit(rune(input[0])) {
		return "", errMsg
	}

	var builder strings.Builder
	var lastChar rune
	charSet := false
	escaped := false

	for _, r := range input {
		if escaped {
			if r == '\\' || unicode.IsDigit(r) {
				builder.WriteRune(r)
				escaped = false
			} else {
				return "", errMsg
			}
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsLetter(r) {
			lastChar = r
			builder.WriteRune(r)
			charSet = true
		} else if unicode.IsDigit(r) {
			if !charSet {
				return "", errMsg
			}
			count, _ := strconv.Atoi(string(r))
			if count == 0 {
				result := builder.String()
				builder.Reset()
				builder.WriteString(result[:len(result)-1])
			} else {
				builder.WriteString(strings.Repeat(string(lastChar), count-1))
			}
		} else {
			return "", errMsg
		}
	}

	if escaped {
		return "", errMsg
	}

	return builder.String(), nil
}
