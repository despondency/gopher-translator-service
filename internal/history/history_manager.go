package history

import "context"

type Manager interface {
	Add(ctx context.Context, word, translation string)
	GetTranslationHistory(ctx context.Context) []*Entry
}
