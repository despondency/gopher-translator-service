package translator

import (
	"fmt"
	"strings"
	"unicode"
)

type StorageManager interface {
	Store(word, translated string)
	Get(word string) (string, bool)
}

type TranslatedWord struct {
	Word        string
	Translation string
}

type HistoryManager interface {
	Add(word, translation string)
	Fetch() []*TranslatedWord
}

var ErrShortenedWord = fmt.Errorf("cannot understand words with '")
var ErrContainsDigits = fmt.Errorf("cannot understand words with digits")
var ErrEmptyWord = fmt.Errorf("cannot translate empty words")

type gopherTranslator struct {
	translatorRules []*translatorRule
	historySvc      HistoryManager
}

func (t *gopherTranslator) TranslateSentence(sentence []string) ([]string, error) {
	translatedWords := make([]string, len(sentence))
	for i, word := range sentence {
		currentTranslatedWord, err := t.Translate(word)
		if err != nil {
			return nil, err
		}
		translatedWords[i] = currentTranslatedWord
	}
	return translatedWords, nil
}

func (t *gopherTranslator) GetTranslateHistory() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

type translatorRule struct {
	Apply func(word string) (string, bool)
}

func NewGopherTranslator(historySvc HistoryManager) Translator {
	return &gopherTranslator{
		historySvc:      historySvc,
		translatorRules: createTranslatorRules(),
	}
}

func (t *gopherTranslator) Translate(word string) (string, error) {
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
				t.historySvc.Add(word, translatedWord)
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
			if strings.HasPrefix(word, "xr") {
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
