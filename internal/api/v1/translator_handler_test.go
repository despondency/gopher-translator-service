package v1

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	translatormock "gopher-translator-service/internal/translator/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTranslateWord(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/word", strings.NewReader(`{"english_word":"apple"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := translatormock.NewGopherTranslator(ctrl)
	m.EXPECT().Translate(context.Background(), "apple").Return("gapple", nil)

	h := NewTranslatorHandler(m)
	handler := h.TranslateWord()
	expectedJSON := `{"gopher_word":"gapple"}`
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJSON+"\n", rec.Body.String())
	}
}
