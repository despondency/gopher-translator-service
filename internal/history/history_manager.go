package history

import "context"

//go:generate mockgen -source=history_manager.go -package=historymock -destination=./mocks/history_manager_mock.go -mock_names=Manager=HistoryManager
type Manager interface {
	Add(ctx context.Context, word, translation string)
	GetTranslationHistory(ctx context.Context) []*Entry
}
