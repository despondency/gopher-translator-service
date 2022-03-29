package translator

import "context"

type Manager interface {
	Translate(ctx context.Context, word string) (string, error)
	TranslateSentence(ctx context.Context, sentence string) (string, error)
}
