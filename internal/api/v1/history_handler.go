package v1

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gopher-translator-service/internal/history"
	"net/http"
	"sort"
)

type TranslationHistory struct {
	History []*history.Entry
}

func (omap *TranslationHistory) UnmarshalJSON(b []byte) error {
	anonymousMap := struct {
		History map[string]string `json:"history"`
	}{}
	err := json.Unmarshal(b, &anonymousMap)
	if err != nil {
		return err
	}
	for key, value := range anonymousMap.History {
		omap.History = append(omap.History, &history.Entry{Word: key, Translation: value})
	}
	sort.Slice(omap.History, func(i, j int) bool {
		return omap.History[i].Word < omap.History[j].Word
	})
	return nil
}

// MarshalJSON Implement the json.Marshaler interface
func (omap TranslationHistory) MarshalJSON() ([]byte, error) {
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
		translationHistory := hh.h.GetTranslationHistory(c.Request().Context())
		resp := &TranslationHistory{History: translationHistory}
		return c.JSON(http.StatusOK, &resp)
	}
}
