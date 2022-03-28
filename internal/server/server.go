package server

import (
	"github.com/labstack/echo/v4"
	v1 "gopher-translator-service/internal/api/v1"
	"gopher-translator-service/internal/translator"
	"log"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	e := echo.New()
	translator := translator.NewGopherTranslator(translator.NewHistoryService())
	translatorHandler := v1.NewTranslatorHandler(translator)
	e.Add(echo.GET, "/health", createHealth())
	e.Add(echo.POST, "/word", translatorHandler.TranslateWord())

	log.Fatalln(e.Start(":8080"))
}
