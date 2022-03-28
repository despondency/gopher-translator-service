package storage

type Storage interface {
	Store(word, translated string)
	Get(word string) (string, bool)
}
