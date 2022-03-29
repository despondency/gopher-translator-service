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
		input, expected string
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
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.Translate(ctx, tt.input)
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
		input, expected string
	}{
		{"apple ear oak user xray chair square xxxxxxxxqqu aaaaplequ!",
			"gapple gear goak guser gexray airchogo aresquogo xxxxxxxxqquogo gaaaaplequ!"},
		{"Apples grow on trees.",
			"gApples owgrogo gon eestrogo."},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation, err := translator.TranslateSentence(ctx, tt.input)
			if err != nil {
				t.Errorf(err.Error())
				t.FailNow()
			}
			if translation != tt.expected {
				t.Errorf("expected %s, got %s", translation, tt.expected)
			}
		})
	}
}

func TestGopherTranslatorTranslateErrors(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
	var tests = []struct {
		input    string
		expected error
	}{
		{"don't", ErrShortenedWord},
		{"shouldn't", ErrShortenedWord},
		{"5asdfhg61261611", ErrContainsDigits},
		{"", ErrEmptyWord},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			_, err := translator.Translate(ctx, tt.input)
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
		input    string
		expected error
	}{
		{"heeeey its me", ErrInvalidSentence},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			_, err := translator.TranslateSentence(ctx, tt.input)
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
