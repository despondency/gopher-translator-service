package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	v1 "gopher-translator-service/internal/api/v1"
	"gopher-translator-service/internal/history"
	"gopher-translator-service/internal/translator"
	"log"
)

type Server struct {
	e    *echo.Echo
	port int
}

func NewServer(port int) *Server {
	return &Server{
		e:    echo.New(),
		port: port,
	}
}

func (s *Server) Run() {
	historySvc := history.NewHistoryService()
	translatorSvc := translator.NewGopherTranslator(historySvc)
	translatorHandler := v1.NewTranslatorHandler(translatorSvc)
	historyHandler := v1.NewHistoryHandler(historySvc)
	s.e.Add(echo.GET, "/v1/health", createHealth())
	s.e.Add(echo.POST, "/v1/word", translatorHandler.TranslateWord())
	s.e.Add(echo.POST, "/v1/sentence", translatorHandler.TranslateSentence())
	s.e.Add(echo.GET, "/v1/history", historyHandler.GetTranslationHistory())
	log.Fatalln(s.e.Start(fmt.Sprintf(":%d", s.port)))
}

func (s *Server) Stop(ctx context.Context) {
	err := s.e.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
