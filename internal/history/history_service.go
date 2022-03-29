package history

import (
	"context"
	"sort"
	"sync"
)

type historyService struct {
	hist map[string]string
	l    *sync.RWMutex
}

type Entry struct {
	Word        string
	Translation string
}

func NewHistoryService() *historyService {
	return &historyService{
		hist: make(map[string]string),
		l:    &sync.RWMutex{},
	}
}

func (hs *historyService) Add(ctx context.Context, word, translation string) {
	hs.l.Lock()
	defer hs.l.Unlock()
	hs.hist[word] = translation
}

func (hs *historyService) GetTranslationHistory(ctx context.Context) []*Entry {
	hs.l.RLock()
	defer hs.l.RUnlock()
	keys := make([]string, 0, len(hs.hist))
	for k, _ := range hs.hist {
		keys = append(keys, k)
	}
	// let's agree that in a real life scenario
	// we will be putting them in a RBTree/AVLTree
	// and won't be sorting on every call by wasting CPU
	sort.Strings(keys)
	wordPair := make([]*Entry, len(keys))
	for i, k := range keys {
		wordPair[i] = &Entry{
			Word:        k,
			Translation: hs.hist[k],
		}
	}
	return wordPair
}
