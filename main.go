package main

import (
	"errors"
	"flag"
	"log/slog"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":80", "server port")
}

func main() {
	flag.Parse()
	logger := slog.Default()
	repo := Repository{
		Logger: logger,
	}
	handler := Handler{
		Logger:     logger,
		Repository: repo,
	}
	handler.Register(http.DefaultServeMux)
	// demo 工程，未加 graceful shutdown
	logger.Info("服务启动", slog.String("addr", addr))
	err := http.ListenAndServe(addr, nil)
	if !errors.Is(err, http.ErrServerClosed) {
		logger.Error("服务启动失败", slog.Any("err", err))
	}
}
