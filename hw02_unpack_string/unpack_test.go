package hw02unpackstring

import (
	"errors"
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "a0b0c0", expected: ""},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "a1b1c1", expected: "abc"},
		{input: `\\`, expected: `\`},
		{input: `a\0b`, expected: `a0b`},
		{input: `\9`, expected: `9`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidMultiDigitCount(t *testing.T) {
	invalidInputs := []string{
		"a10",
		"aaa10b",
	}
	for _, input := range invalidInputs {
		input := input
		t.Run(input, func(t *testing.T) {
			_, err := Unpack(input)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidEscaping(t *testing.T) {
	invalidInputs := []string{
		`qw\ne`,
		`\`,
		`a\`,
		`\q`,
		`qwe\a5`,
		`qwe\\\\q`,
	}

	for _, input := range invalidInputs {
		input := input
		t.Run(input, func(t *testing.T) {
			_, err := Unpack(input)
			require.Truef(t, errors.Is(err, ErrInvalidString), "ожидали ошибку при некорректном экранировании: %q", err)
		})
	}
}

func TestUnpackUnicode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "😀2🙂3", expected: "😀😀🙂🙂🙂"},
		{input: "日2本3語1", expected: "日日本本本語"},
		{input: `😀\2`, expected: `😀2`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackSequentialDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a2b3c4", expected: "aabbbcccc"},
		{input: "z9y8x7", expected: "zzzzzzzzzyyyyyyyyxxxxxxx"},
		{input: "a2b0c3", expected: "aaccc"},
		{input: "a0b0c0", expected: ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
