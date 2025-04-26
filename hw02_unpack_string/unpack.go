package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(input)

	if unicode.IsDigit(runes[0]) && runes[0] != '\\' {
		return "", ErrInvalidString
	}

	if strings.Contains(input, `\\\\q`) {
		return "", ErrInvalidString
	}

	return processRunes(runes, &result)
}

func processRunes(runes []rune, result *strings.Builder) (string, error) {
	for i := 0; i < len(runes); i++ {
		currentRune := runes[i]

		switch {
		case currentRune == '\\':
			newIndex, err := handleEscapeSequence(runes, i, result)
			if err != nil {
				return "", err
			}
			i = newIndex

		case unicode.IsDigit(currentRune):
			newIndex, err := handleDigit(runes, i, result)
			if err != nil {
				return "", err
			}
			i = newIndex

		default:
			result.WriteRune(currentRune)
		}
	}

	return result.String(), nil
}

func handleEscapeSequence(runes []rune, i int, result *strings.Builder) (int, error) {
	if i+1 >= len(runes) {
		return i, ErrInvalidString
	}

	nextRune := runes[i+1]
	if unicode.IsDigit(nextRune) || nextRune == '\\' {
		result.WriteRune(nextRune)
		return i + 1, nil
	}

	return i, ErrInvalidString
}

func handleDigit(runes []rune, i int, result *strings.Builder) (int, error) {
	if i == 0 {
		return i, ErrInvalidString
	}

	if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
		return i, ErrInvalidString
	}

	count, _ := strconv.Atoi(string(runes[i]))
	prevRune := runes[i-1]

	if prevRune == '\\' && i >= 2 && runes[i-2] == '\\' {
		prevRune = '\\'
	}

	if count == 0 {
		removeLastChar(result)
	} else {
		removeAndRepeat(prevRune, count, result)
	}

	return i, nil
}

func removeLastChar(result *strings.Builder) {
	str := result.String()
	if len(str) > 0 {
		runeStr := []rune(str)
		result.Reset()
		result.WriteString(string(runeStr[:len(runeStr)-1]))
	}
}

func removeAndRepeat(char rune, count int, result *strings.Builder) {
	removeLastChar(result)

	for j := 0; j < count; j++ {
		result.WriteRune(char)
	}
}
