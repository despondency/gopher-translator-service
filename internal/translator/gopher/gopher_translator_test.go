package gopher

import (
	"fmt"
	"testing"
)

func TestGopherTranslator(t *testing.T) {
	translator := NewGopherTranslator()
	var tests = []struct {
		actual, expected string
	}{
		{"apple", "gapple"},
		{"ear", "gear"},
		{"oak", "goak"},
		{"user", "guser"},
		{"xray", "gexray"},
		{"chair", "airchogo"},
		{"square", "aresquogo"},
		{"xxxxxxxxqqu", "xxxxxxxxqquogo"},
		{"aaaaplequ", "gaaaaplequ"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.Translate(tt.actual)
			if err != nil {
				t.Errorf(err.Error())
				t.FailNow()
			}
			if translation != tt.expected {
				t.Errorf("got %s, want %s", translation, tt.expected)
			}
		})
	}
}

func TestGopherTranslatorErrors(t *testing.T) {
	translator := NewGopherTranslator()
	var tests = []struct {
		actual   string
		expected error
	}{
		{"don't", ErrShortenedWord},
		{"shouldn't", ErrShortenedWord},
		{"5asdfhg61261611", ErrContainsDigits},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			_, err := translator.Translate(tt.actual)
			if err == nil {
				t.Errorf("expects an error to occur")
				t.FailNow()
			}
			if err != tt.expected {
				t.Errorf("got %s, want %s", err, tt.expected)
			}
		})
	}
}
