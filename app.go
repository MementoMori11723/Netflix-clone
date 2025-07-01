package main

import (
	"log/slog"
	"net/http"
	"netflix-clone/server"
	"os"
)

func main() {
	_mux := server.New()
	_server := http.Server{
		Addr:    ":8080",
		Handler: _mux,
	}
	slog.Info(
		"Server Running...",
		"URL", "http://localhost"+_server.Addr,
	)
	if err := _server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
