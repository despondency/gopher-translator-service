package translator

import (
	"fmt"
	"testing"
)

func TestGopherTranslator_Translate(t *testing.T) {
	translator := NewGopherTranslator(NewHistoryService())
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

func TestGopherTranslator_TranslateSentence(t *testing.T) {
	translator := NewGopherTranslator(NewHistoryService())
	var tests = []struct {
		actual, expected []string
	}{
		{[]string{"apple", "ear", "oak", "user", "xray", "chair", "square", "xxxxxxxxqqu", "aaaaplequ"},
			[]string{"gapple", "gear", "goak", "guser", "gexray", "airchogo", "aresquogo", "xxxxxxxxqquogo", "gaaaaplequ"}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.TranslateSentence(tt.actual)
			if err != nil {
				t.Errorf(err.Error())
				t.FailNow()
			}
			compareSentences(t, translation, tt.expected)
		})
	}
}

func compareSentences(t *testing.T, actual []string, expected []string) {
	if len(actual) != len(expected) {
		t.Errorf("len of the expected sentances and actual is not equal, expected is %d, actual is %d", len(expected), len(actual))
	}
	for i, actualWord := range actual {
		if actualWord != expected[i] {
			t.Errorf("actual is %s, expected is %s", actualWord, expected[i])
		}
	}
}

func TestGopherTranslatorErrors(t *testing.T) {
	translator := NewGopherTranslator(NewHistoryService())
	var tests = []struct {
		actual   string
		expected error
	}{
		{"don't", ErrShortenedWord},
		{"shouldn't", ErrShortenedWord},
		{"5asdfhg61261611", ErrContainsDigits},
		{"", ErrEmptyWord},
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
