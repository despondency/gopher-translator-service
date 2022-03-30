package translator

import "context"

//go:generate mockgen -source=translator_manager.go -package=translatormock -destination=./mocks/translator_manager_mock.go -mock_names=Manager=GopherTranslator
type Manager interface {
	Translate(ctx context.Context, word string) string
	TranslateSentence(ctx context.Context, sentence string) string
}
