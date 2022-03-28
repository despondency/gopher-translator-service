package main

import "gopher-translator-service/internal/server"

func main() {
	srv := server.NewServer()
	srv.Run()
}
