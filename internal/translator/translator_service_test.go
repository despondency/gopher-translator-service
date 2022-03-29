package translator

import (
	"context"
	"fmt"
	"gopher-translator-service/internal/history"
	"testing"
)

func TestGopherTranslatorTranslate(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
	var tests = []struct {
		actual, expected string
	}{
		{"apple", "gapple"},
		{"Apple", "gApple"},
		{"ear", "gear"},
		{"Ear", "gEar"},
		{"oak", "goak"},
		{"Oak", "gOak"},
		{"user", "guser"},
		{"User", "gUser"},
		{"xray", "gexray"},
		{"XRay", "geXRay"},
		{"chair", "airchogo"},
		{"Chair", "airChogo"},
		{"square", "aresquogo"},
		{"Square", "areSquogo"},
		{"xxxxxxxxqqu", "xxxxxxxxqquogo"},
		{"Xxxxxxxxqqu", "Xxxxxxxxqquogo"},
		{"aaaaplequ", "gaaaaplequ"},
		{"Aaaaplequ", "gAaaaplequ"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.Translate(ctx, tt.actual)
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

func TestGopherTranslatorTranslateSentence(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
	var tests = []struct {
		actual, expected string
	}{
		{"apple ear oak user xray chair square xxxxxxxxqqu aaaaplequ!",
			"gapple gear goak guser gexray airchogo aresquogo xxxxxxxxqquogo gaaaaplequ!"},
		{"Apples grow on trees.",
			"gApples owgrogo gon eestrogo."},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.TranslateSentence(ctx, tt.actual)
			if err != nil {
				t.Errorf(err.Error())
				t.FailNow()
			}
			if translation != tt.expected {
				t.Errorf("expected %s, got %s", translation, tt.expected)
			}
			//compareSentences(t, translation, tt.expected)
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

func TestGopherTranslatorTranslateErrors(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
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
			_, err := translator.Translate(ctx, tt.actual)
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

func TestGopherTranslatorTranslateSentenceErrors(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
	var tests = []struct {
		actual   string
		expected error
	}{
		{"heeeey its me", ErrInvalidSentence},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.actual, tt.expected)
		t.Run(testName, func(t *testing.T) {
			_, err := translator.TranslateSentence(ctx, tt.actual)
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
