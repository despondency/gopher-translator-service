package translator

type Translator interface {
	Translate(word string) (string, error)
}
