package history

import (
	"context"
	"testing"
)

func TestHistoryService(t *testing.T) {
	historySvc := NewHistoryService()
	ctx := context.Background()
	var toAdd = []struct {
		word, translation string
	}{
		{"c", "h"},
		{"b", "h"},
		{"a", "h"},
	}
	for _, curr := range toAdd {
		historySvc.Add(ctx, curr.word, curr.translation)
	}
	translatedWordsActual := historySvc.GetTranslationHistory(ctx)
	compareHistory(t, translatedWordsActual, []*Entry{
		{
			Word:        "a",
			Translation: "h",
		},
		{
			Word:        "b",
			Translation: "h",
		},
		{
			Word:        "c",
			Translation: "h",
		},
	})
}

func compareHistory(t *testing.T, actual []*Entry, expected []*Entry) {
	if len(actual) != len(expected) {
		t.Errorf("got %d len, want %d len", len(actual), len(expected))
	}
	for i, act := range actual {
		if act.Word != expected[i].Word {
			t.Errorf("got %s, want %s Word", act.Word, expected[i].Word)
		}
		if act.Translation != expected[i].Translation {
			t.Errorf("got %s, want %s translation", act.Translation, expected[i].Translation)
		}
	}
}
