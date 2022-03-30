package v1

import (
	"fmt"
	"testing"
)

func TestGopherTranslator_TranslateErrors(t *testing.T) {
	tr := NewTranslatorRequestValidator()
	var tests = []struct {
		input    *GopherWordRequest
		expected error
	}{
		{&GopherWordRequest{EnglishWord: "don't"}, ErrShortened},
		{&GopherWordRequest{EnglishWord: "shouldn't"}, ErrShortened},
		{&GopherWordRequest{EnglishWord: "5asdfhg61261611"}, ErrContainsDigits},
		{&GopherWordRequest{EnglishWord: ""}, ErrEmpty},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			err := tr.ValidateWordReq(tt.input)
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

func TestGopherTranslator_TranslateSentenceErrors(t *testing.T) {
	tr := NewTranslatorRequestValidator()
	var tests = []struct {
		input    *GopherSentenceRequest
		expected error
	}{
		{&GopherSentenceRequest{EnglishSentence: ""}, ErrEmpty},
		{&GopherSentenceRequest{EnglishSentence: "heeeey its me"}, ErrInvalidSentence},
		{&GopherSentenceRequest{EnglishSentence: "ė"}, ErrNotEnglish},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			err := tr.ValidateSentenceReq(tt.input)
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
