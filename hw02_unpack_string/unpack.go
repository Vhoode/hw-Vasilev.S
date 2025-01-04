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

	for i := 0; i < len(input); i++ {
		r := rune(input[i])

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

		switch {
		case unicode.IsLetter(r):
			lastChar = r
			builder.WriteRune(r)
			charSet = true
		case unicode.IsDigit(r):
			if !charSet {
				return "", errMsg
			}

			countStr := string(r)
			for i+1 < len(input) && unicode.IsDigit(rune(input[i+1])) {
				countStr += string(input[i+1])
				i++
			}

			count, err := strconv.Atoi(countStr)
			if err != nil || count > 9 {
				return "", errMsg
			}

			if count == 0 {
				result := builder.String()
				builder.Reset()
				builder.WriteString(result[:len(result)-1])
			} else {
				builder.WriteString(strings.Repeat(string(lastChar), count-1))
			}
		default:
			return "", errMsg
		}
	}

	if escaped {
		return "", errMsg
	}

	return builder.String(), nil
}
