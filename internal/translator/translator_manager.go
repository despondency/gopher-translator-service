package translator

import "context"

//go:generate mockgen -destination=./mocks/mock_doer.go -package=mocks Manager
type Manager interface {
	Translate(ctx context.Context, word string) (string, error)
	TranslateSentence(ctx context.Context, sentence string) (string, error)
}
