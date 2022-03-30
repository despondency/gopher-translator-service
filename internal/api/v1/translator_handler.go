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
	t         translator.Manager
	validator *TranslatorRequestValidator
}

func NewTranslatorHandler(t translator.Manager) *TranslatorHandler {
	return &TranslatorHandler{
		t:         t,
		validator: NewTranslatorRequestValidator(),
	}
}

func (th *TranslatorHandler) TranslateWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GopherWordRequest
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		if err := th.validator.ValidateWordReq(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation := th.t.Translate(c.Request().Context(), req.EnglishWord)
		return c.JSON(http.StatusOK, &GopherWordResponse{GopherWord: translation})
	}
}

func (th *TranslatorHandler) TranslateSentence() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GopherSentenceRequest
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		if err := th.validator.ValidateSentenceReq(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation := th.t.TranslateSentence(c.Request().Context(), req.EnglishSentence)
		return c.JSON(http.StatusOK, &GopherSentenceResponse{GopherSentence: translation})
	}
}
