package v1

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	translatormock "gopher-translator-service/internal/translator/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_TranslateWord(t *testing.T) {
	var tests = []struct {
		mockInput  string
		mockOutput string
		mockFunc   func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator
		statusCode int
		inputJson  string
		resultJson string
	}{
		{
			mockInput: "apple", mockOutput: "gapple",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				m := translatormock.NewGopherTranslator(ctrl)
				m.EXPECT().Translate(context.Background(), input).Return(output).Times(1)
				return m
			},
			statusCode: 200,
			resultJson: `{"gopher_word":"gapple"}` + "\n",
			inputJson:  `{"english_word":"apple"}`,
		},
		{
			mockInput: "", mockOutput: "",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				return nil
			},
			statusCode: 400,
			resultJson: ``,
			inputJson:  `{"english_word":""}`,
		},
		{
			mockInput: "", mockOutput: "",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				return nil
			},
			statusCode: 400,
			resultJson: ``,
			inputJson:  `"english_word":"apple"}`,
		},
	}
	e := echo.New()
	ctrl := gomock.NewController(t)
	for i, tt := range tests {
		testName := fmt.Sprintf("testTranslateWord-%d", i)
		t.Run(testName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/word", strings.NewReader(tt.inputJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			m := tt.mockFunc(ctrl, tt.mockInput, tt.mockOutput)
			h := NewTranslatorHandler(m)
			handler := h.TranslateWord()
			if assert.NoError(t, handler(c)) {
				assert.Equal(t, tt.statusCode, rec.Code)
				assert.Equal(t, tt.resultJson, rec.Body.String())
			}
			ctrl.Finish()
		})
	}
}

func Test_TranslateSentence(t *testing.T) {
	var tests = []struct {
		mockInput  string
		mockOutput string
		mockFunc   func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator
		statusCode int
		inputJson  string
		resultJson string
	}{
		{
			mockInput: "Apples grow on trees.", mockOutput: "gApples owgrogo gon eestrogo.",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				m := translatormock.NewGopherTranslator(ctrl)
				m.EXPECT().TranslateSentence(context.Background(), input).Return(output).Times(1)
				return m
			},
			statusCode: 200,
			resultJson: `{"gopher_sentence":"gApples owgrogo gon eestrogo."}` + "\n",
			inputJson:  `{"english_sentence":"Apples grow on trees."}`,
		},
		{
			mockInput: "Apples grow on trees", mockOutput: "",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				return nil
			},
			statusCode: 400,
			resultJson: ``,
			inputJson:  `{"english_sentence":"Apples grow on trees"}`,
		},
		{
			mockInput: "", mockOutput: "",
			mockFunc: func(ctrl *gomock.Controller, input string, output string) *translatormock.GopherTranslator {
				return nil
			},
			statusCode: 400,
			resultJson: ``,
			inputJson:  `"english_sentence":"Apples grow on trees."}`,
		},
	}
	e := echo.New()
	ctrl := gomock.NewController(t)
	for i, tt := range tests {
		testName := fmt.Sprintf("testTranslateSentence-%d", i)
		t.Run(testName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/sentence", strings.NewReader(tt.inputJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			m := tt.mockFunc(ctrl, tt.mockInput, tt.mockOutput)
			h := NewTranslatorHandler(m)
			handler := h.TranslateSentence()
			if assert.NoError(t, handler(c)) {
				assert.Equal(t, tt.statusCode, rec.Code)
				assert.Equal(t, tt.resultJson, rec.Body.String())
			}
			ctrl.Finish()
		})
	}
}
