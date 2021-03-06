package translator

import (
	"context"
	"gopher-translator-service/internal/history"
	"strings"
)

type gopherTranslator struct {
	translatorRules []*translatorRule
	historySvc      history.Manager
}

func (t *gopherTranslator) TranslateSentence(ctx context.Context, sentence string) string {
	endsWith := []rune(sentence)
	end := endsWith[len(sentence)-1]
	newSentence := strings.ReplaceAll(sentence, string(end), "")
	sentenceWords := strings.Fields(newSentence)
	translatedWords := make([]string, len(sentenceWords))
	commas := make([]bool, len(sentenceWords))
	for i, word := range sentenceWords {
		if strings.HasSuffix(word, ",") {
			commas[i] = true
			sentenceWords[i] = strings.ReplaceAll(sentenceWords[i], ",", "")
		}
	}
	for i, word := range sentenceWords {
		currentTranslatedWord := t.translate(word)
		translatedWords[i] = currentTranslatedWord
	}
	translationStrBuilder := strings.Builder{}
	translation := strings.Builder{}
	for i, word := range translatedWords {
		if commas[i] {
			translation.WriteString(word)
			translation.WriteString(", ")
		} else {
			translation.WriteString(word)
			if i+1 < len(translatedWords) {
				translation.WriteString(" ")
			}
		}
	}
	translationStrBuilder.WriteString(translation.String())
	translationStrBuilder.WriteRune(end)
	t.historySvc.Add(ctx, sentence, translationStrBuilder.String())
	return translationStrBuilder.String()
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

func (t *gopherTranslator) Translate(ctx context.Context, word string) string {
	translation := t.translate(word)
	t.historySvc.Add(ctx, word, translation)
	return translation
}

func (t *gopherTranslator) translate(word string) string {
	for _, rule := range t.translatorRules {
		translatedWord, applied := rule.Apply(word)
		if applied {
			return translatedWord
		}
	}
	return word
}

func createTranslatorRules() []*translatorRule {
	var vowels = map[string]struct{}{
		"a": {}, "e": {}, "i": {}, "o": {}, "u": {},
		"A": {}, "E": {}, "I": {}, "O": {}, "U": {},
	}
	rules := make([]*translatorRule, 3)
	rules[0] = &translatorRule{
		Apply: func(word string) (string, bool) {
			if _, ok := vowels[string(word[0])]; ok {
				builder := strings.Builder{}
				// starts with vowel
				builder.WriteString("g")
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
			if _, ok := vowels[string(word[0])]; !ok {
				builder := strings.Builder{}
				// starts with consonant
				var end = 0
				// get end of the consonant sound
				for ; end < len(word); end++ {
					if _, ok = vowels[string(word[end])]; ok {
						break
					}
				}
				// check if the last match is a 'q', check next if its vowel 'u' to get special 'qu'
				if end > 1 && word[end-1] == 'q' && word[end] == 'u' {
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
