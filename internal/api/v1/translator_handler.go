package v1

import (
	"bytes"
	"encoding/json"
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

// MarshalJSON Implement the json.Marshaler interface
func (omap GetTranslationHistory) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\"history\":{")
	for i, kv := range omap.History {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(kv.Word)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value
		val, err := json.Marshal(kv.Translation)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}}")
	return buf.Bytes(), nil
}

type TranslatorHandler struct {
	t translator.Manager
}

func NewTranslatorHandler(t translator.Manager) *TranslatorHandler {
	return &TranslatorHandler{
		t: t,
	}
}

func (th *TranslatorHandler) TranslateWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GopherWordRequest
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation, err := th.t.Translate(c.Request().Context(), req.EnglishWord)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		var gopherWordResponse GopherWordResponse
		gopherWordResponse.GopherWord = translation
		return c.JSON(http.StatusOK, &gopherWordResponse)
	}
}

func (th *TranslatorHandler) TranslateSentence() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GopherSentenceRequest
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation, err := th.t.TranslateSentence(c.Request().Context(), req.EnglishSentence)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		var gopherWordResponse GopherWordResponse
		gopherWordResponse.GopherWord = translation
		return c.JSON(http.StatusOK, &gopherWordResponse)
	}
}
