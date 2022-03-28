package translator

type Translator interface {
	Translate(word string) (string, error)
	TranslateSentence(sentence []string) ([]string, error)
	GetTranslateHistory() ([]string, error)
}
