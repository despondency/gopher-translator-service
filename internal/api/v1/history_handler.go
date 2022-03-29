package v1

import (
	"github.com/labstack/echo/v4"
	"gopher-translator-service/internal/history"
	"net/http"
)

type GetTranslationHistory struct {
	History []*history.Entry
}

type HistoryHandler struct {
	h history.Manager
}

func NewHistoryHandler(h history.Manager) *HistoryHandler {
	return &HistoryHandler{
		h: h,
	}
}

func (hh *HistoryHandler) GetTranslationHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GopherSentenceRequest
		if err := c.Bind(&req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		translation := hh.h.GetTranslationHistory(c.Request().Context())
		var resp GetTranslationHistory
		resp.History = translation
		return c.JSON(http.StatusOK, &resp)
	}
}
