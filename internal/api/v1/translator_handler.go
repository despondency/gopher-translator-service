package v1

import (
	"github.com/labstack/echo/v4"
	"gopher-translator-service/internal/translator"
	"net/http"
)

type GopherSentenceRequest struct {
	EnglishSentence string `json:"english_sentence"`
}

type GopherSentenceResponse struct {
	GopherSentence string `json:"gopher_sentence"`
}

type GopherWordRequest struct {
	EnglishWord string `json:"english_word"`
}

type GopherWordResponse struct {
	GopherWord string `json:"gopher_word"`
}

type TranslatorHandler struct {
	t translator.Translator
}

func NewTranslatorHandler(t translator.Translator) *TranslatorHandler {
	return &TranslatorHandler{
		t: t,
	}
}

func (th *TranslatorHandler) TranslateWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		var gopherWordRequest GopherWordRequest
		if err := c.Bind(&gopherWordRequest); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation, err := th.t.Translate(gopherWordRequest.EnglishWord)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		var gopherWordResponse GopherWordResponse
		gopherWordResponse.GopherWord = translation
		return c.JSON(http.StatusOK, &gopherWordResponse)
	}
}
