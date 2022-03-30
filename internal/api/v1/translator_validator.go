package v1

import (
	"fmt"
	"strings"
	"unicode"
)

var ErrShortened = fmt.Errorf("cannot understand words with '")
var ErrContainsDigits = fmt.Errorf("cannot understand words with digits")
var ErrEmpty = fmt.Errorf("cannot translate empty words")
var ErrInvalidSentence = fmt.Errorf("sentence does not end in (.?!)")

type TranslatorRequestValidator struct {
}

func NewTranslatorRequestValidator() *TranslatorRequestValidator {
	return &TranslatorRequestValidator{}
}

func (trv *TranslatorRequestValidator) ValidateWordReq(request *GopherWordRequest) error {
	return trv.validate(request.EnglishWord)
}

func (trv *TranslatorRequestValidator) ValidateSentenceReq(request *GopherSentenceRequest) error {
	sentence := request.EnglishSentence
	err := trv.validate(sentence)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(sentence, "!") && !strings.HasSuffix(sentence, "?") && !strings.HasSuffix(sentence, ".") {
		return ErrInvalidSentence
	}
	return nil
}

func (trv *TranslatorRequestValidator) validate(s string) error {
	if s == "" {
		return ErrEmpty
	}
	if strings.Contains(s, "'") {
		return ErrShortened
	}
	if containsNumber(s) {
		return ErrContainsDigits
	}
	return nil
}

func containsNumber(word string) bool {
	runes := []rune(word)
	for _, r := range runes {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
