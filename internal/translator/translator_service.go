package translator

import (
	"context"
	"fmt"
	"gopher-translator-service/internal/history"
	"strings"
	"unicode"
)

var ErrShortenedWord = fmt.Errorf("cannot understand words with '")
var ErrContainsDigits = fmt.Errorf("cannot understand words with digits")
var ErrEmptyWord = fmt.Errorf("cannot translate empty words")
var ErrInvalidSentence = fmt.Errorf("sentence does not end in (.?!)")

type gopherTranslator struct {
	translatorRules []*translatorRule
	historySvc      history.Manager
}

func (t *gopherTranslator) TranslateSentence(ctx context.Context, sentence string) (string, error) {
	if !strings.HasSuffix(sentence, "!") && !strings.HasSuffix(sentence, "?") && !strings.HasSuffix(sentence, ".") {
		return "", ErrInvalidSentence
	}
	endsWith := []rune(sentence)
	end := endsWith[len(sentence)-1]
	newSentence := strings.ReplaceAll(sentence, string(end), "")
	sentenceWords := strings.Fields(newSentence)
	translatedWords := make([]string, len(sentenceWords))
	for i, word := range sentenceWords {
		currentTranslatedWord, err := t.translate(word)
		if err != nil {
			return "", err
		}
		translatedWords[i] = currentTranslatedWord
	}
	translationStrBuilder := strings.Builder{}
	translation := strings.Join(translatedWords, " ")
	translationStrBuilder.WriteString(translation)
	translationStrBuilder.WriteRune(end)
	t.historySvc.Add(ctx, sentence, translationStrBuilder.String())
	return translationStrBuilder.String(), nil
}

type translatorRule struct {
	Apply func(word string) (string, bool)
}

func NewGopherTranslator(historySvc history.Manager) Manager {
	return &gopherTranslator{
		historySvc:      historySvc,
		translatorRules: createTranslatorRules(),
	}
}

func (t *gopherTranslator) Translate(ctx context.Context, word string) (string, error) {
	translation, err := t.translate(word)
	if err != nil {
		return "", err
	}
	t.historySvc.Add(ctx, word, translation)
	return translation, nil
}

func (t *gopherTranslator) translate(word string) (string, error) {
	if len(word) > 0 {
		if strings.Contains(word, "'") {
			return word, ErrShortenedWord
		}
		if containsNumber(word) {
			return word, ErrContainsDigits
		}
		for _, rule := range t.translatorRules {
			translatedWord, applied := rule.Apply(word)
			if applied {
				return translatedWord, nil
			}
		}
	}
	return word, ErrEmptyWord
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

func createTranslatorRules() []*translatorRule {
	var vowels = map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	rules := make([]*translatorRule, 3)
	rules[0] = &translatorRule{
		Apply: func(word string) (string, bool) {
			runes := []rune(word)
			if _, ok := vowels[runes[0]]; ok {
				builder := strings.Builder{}
				// starts with vowel
				builder.WriteRune('g')
				builder.WriteString(word)
				return builder.String(), true
			}
			return word, false
		},
	}
	rules[1] = &translatorRule{
		Apply: func(word string) (string, bool) {
			builder := strings.Builder{}
			if strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "XR") {
				builder.WriteString("ge")
				builder.WriteString(word)
				return builder.String(), true
			}
			return word, false
		},
	}
	rules[2] = &translatorRule{
		Apply: func(word string) (string, bool) {
			runes := []rune(word)
			if _, ok := vowels[runes[0]]; !ok {
				builder := strings.Builder{}
				// starts with consonant
				var end = 0
				// get end of the consonant sound
				for ; end < len(runes); end++ {
					if _, ok = vowels[runes[end]]; ok {
						break
					}
					end++
				}
				// check if the last match is a 'q', check next if its vowel 'u' to get special 'qu'
				if end > 1 && runes[end-1] == 'q' && runes[end] == 'u' {
					// we have special 'qu'
					end++
					builder.WriteString(word[end:])
					builder.WriteString(word[0:end])
					builder.WriteString("ogo")
					return builder.String(), true
				} else {
					builder.WriteString(word[end:])
					builder.WriteString(word[0:end])
					builder.WriteString("ogo")
					return builder.String(), true
				}
			}
			return word, false
		},
	}
	return rules
}
