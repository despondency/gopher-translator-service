package server

import (
	"context"
	"github.com/labstack/echo/v4"
	v1 "gopher-translator-service/internal/api/v1"
	"gopher-translator-service/internal/history"
	"gopher-translator-service/internal/translator"
	"log"
)

type Server struct {
	e           *echo.Echo
	stopperChan chan struct{}
}

func NewServer() *Server {
	return &Server{
		e:           echo.New(),
		stopperChan: make(chan struct{}),
	}
}

func (s *Server) Run() {
	historySvc := history.NewHistoryService()
	translatorSvc := translator.NewGopherTranslator(history.NewHistoryService())
	translatorHandler := v1.NewTranslatorHandler(translatorSvc)
	historyHandler := v1.NewHistoryHandler(historySvc)
	s.e.Add(echo.GET, "/v1/health", createHealth())
	s.e.Add(echo.POST, "/v1/word", translatorHandler.TranslateWord())
	s.e.Add(echo.POST, "/v1/sentence", translatorHandler.TranslateSentence())
	s.e.Add(echo.GET, "/v1/history", historyHandler.GetTranslationHistory())
	log.Fatalln(s.e.Start(":8080"))
}

func (s *Server) Stop(ctx context.Context) {
	err := s.e.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
