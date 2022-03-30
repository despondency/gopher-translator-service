package v1

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopher-translator-service/internal/history"
	historymock "gopher-translator-service/internal/history/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHistoryHandler_GetTranslationHistory(t *testing.T) {
	var tests = []struct {
		mockOutput []*history.Entry
		mockFunc   func(ctrl *gomock.Controller, output []*history.Entry) *historymock.HistoryManager
		statusCode int
		resultJson string
	}{
		{
			mockOutput: []*history.Entry{
				{"a", "b"},
				{"b", "c"},
			},
			mockFunc: func(ctrl *gomock.Controller, output []*history.Entry) *historymock.HistoryManager {
				m := historymock.NewHistoryManager(ctrl)
				m.EXPECT().GetTranslationHistory(context.Background()).Return(output).Times(1)
				return m
			},
			statusCode: 200,
			resultJson: `{"history":{"a":"b","b":"c"}}` + "\n",
		},
		{
			mockOutput: []*history.Entry{},
			mockFunc: func(ctrl *gomock.Controller, output []*history.Entry) *historymock.HistoryManager {
				m := historymock.NewHistoryManager(ctrl)
				m.EXPECT().GetTranslationHistory(context.Background()).Return(output).Times(1)
				return m
			},
			statusCode: 200,
			resultJson: `{"history":{}}` + "\n",
		},
	}
	e := echo.New()
	ctrl := gomock.NewController(t)
	for i, tt := range tests {
		testName := fmt.Sprintf("testTranslateWord-%d", i)
		t.Run(testName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/history", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			m := tt.mockFunc(ctrl, tt.mockOutput)
			h := NewHistoryHandler(m)
			handler := h.GetTranslationHistory()
			if assert.NoError(t, handler(c)) {
				assert.Equal(t, tt.statusCode, rec.Code)
				assert.Equal(t, tt.resultJson, rec.Body.String())
			}
			ctrl.Finish()
		})
	}
}
