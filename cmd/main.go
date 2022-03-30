package main

import (
	"flag"
	"gopher-translator-service/internal/server"
)

func main() {
	port := flag.Int("port", 8080, "port for starting the application")
	flag.Parse()
	srv := server.NewServer(*port)
	srv.Run()
}
