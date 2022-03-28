package memory

import (
	"gopher-translator-service/internal/translator"
	"sync"
)

type memoryStorage struct {
	store map[string]string
	lock  *sync.RWMutex
}

func NewMemoryStorage() translator.StorageManager {
	return &memoryStorage{
		store: make(map[string]string),
		lock:  &sync.RWMutex{},
	}
}

func (s *memoryStorage) Store(word, translation string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[word] = translation
}

func (s *memoryStorage) Get(word string) (string, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	translation, ok := s.store[word]
	return translation, ok
}
