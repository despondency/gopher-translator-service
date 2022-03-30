package translator

import (
	"context"
	"fmt"
	"gopher-translator-service/internal/history"
	"testing"
)

func TestGopherTranslator_Translate(t *testing.T) {
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
		{"a", "ga"},
		{"c", "cogo"},
		{"e", "ge"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation := translator.Translate(ctx, tt.input)
			if translation != tt.expected {
				t.Errorf("got %s, want %s", translation, tt.expected)
			}
		})
	}
}

func TestGopherTranslator_TranslateSentence(t *testing.T) {
	translator := NewGopherTranslator(history.NewHistoryService())
	ctx := context.Background()
	var tests = []struct {
		input, expected string
	}{
		{"apple ear oak user xray chair square xxxxxxxxqqu aaaaplequ!",
			"gapple gear goak guser gexray airchogo aresquogo xxxxxxxxqquogo gaaaaplequ!"},
		{"Apples grow on trees.",
			"gApples owgrogo gon eestrogo."},
		{"i am sure, that, this is, the one.",
			"gi gam uresogo, atthogo, isthogo gis, ethogo gone."},
		{"I am a true, true legend!",
			"gI gam ga uetrogo, uetrogo egendlogo!"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			translation := translator.TranslateSentence(ctx, tt.input)
			if translation != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, translation)
			}
		})
	}
}
