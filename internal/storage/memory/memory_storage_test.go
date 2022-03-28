package memory

import (
	"fmt"
	"testing"
)

func TestMemoryStorageStore(t *testing.T) {
	store := NewMemoryStorage()
	var tests = []struct {
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
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.word, tt.translation)
		t.Run(testName, func(t *testing.T) {
			go func(currentTest struct {
				word, translation string
			}) {
				store.Store(currentTest.word, currentTest.translation)
				translation, ok := store.Get(currentTest.word)
				if !ok {
					t.Errorf("did not get any response from the storage")
				}
				if translation != currentTest.translation {
					t.Errorf("got %s, want %s", translation, currentTest.translation)
				}
			}(tt)
		})
	}
}
