package translator

import (
	"testing"
)

func TestHistoryService(t *testing.T) {
	historySvc := NewHistoryService()
	var toAdd = []struct {
		word, translation string
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
	for _, curr := range toAdd {
		historySvc.Add(curr.word, curr.translation)
	}
	translatedWordsActual := historySvc.Fetch()
	compareSortedHistory(t, translatedWordsActual, []TranslatedWord{
		{
			Word:        "aaaaplequ",
			Translation: "gaaaaplequ",
		},
		{
			Word:        "apple",
			Translation: "gapple",
		},
		{
			Word:        "chair",
			Translation: "airchogo",
		},
		{
			Word:        "ear",
			Translation: "gear",
		},
		{
			Word:        "oak",
			Translation: "goak",
		},
		{
			Word:        "square",
			Translation: "aresquogo",
		},
		{
			Word:        "user",
			Translation: "guser",
		},
		{
			Word:        "xray",
			Translation: "gexray",
		},
		{
			Word:        "xxxxxxxxqqu",
			Translation: "xxxxxxxxqquogo",
		},
	})
}

func compareSortedHistory(t *testing.T, actual []*TranslatedWord, expected []TranslatedWord) {
	if len(actual) != len(expected) {
		t.Errorf("got %d len, want %d len", len(actual), len(expected))
	}
	for i, act := range actual {
		if act.Word != expected[i].Word {
			t.Errorf("got %s, want %s word", act.Word, expected[i].Word)
		}
		if act.Translation != expected[i].Translation {
			t.Errorf("got %s, want %s translation", act.Translation, expected[i].Translation)
		}
	}
}
