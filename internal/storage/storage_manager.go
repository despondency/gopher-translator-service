package storage

type Manager interface {
	Store(word, translated string)
	Get(word string) (string, bool)
}
