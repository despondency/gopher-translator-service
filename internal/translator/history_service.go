package translator

import (
	"sort"
	"sync"
)

type historyService struct {
	hist map[string]string
	l    *sync.RWMutex
}

func NewHistoryService() *historyService {
	return &historyService{
		hist: make(map[string]string),
		l:    &sync.RWMutex{},
	}
}

func (hs *historyService) Add(word, translation string) {
	hs.l.Lock()
	defer hs.l.Unlock()
	hs.hist[word] = translation
}

func (hs *historyService) Fetch() []*TranslatedWord {
	hs.l.RLock()
	defer hs.l.RUnlock()
	keys := make([]string, 0, len(hs.hist))
	for k, _ := range hs.hist {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	wordPair := make([]*TranslatedWord, len(keys))
	for i, k := range keys {
		wordPair[i] = &TranslatedWord{
			Word:        k,
			Translation: hs.hist[k],
		}
	}
	return wordPair
}
