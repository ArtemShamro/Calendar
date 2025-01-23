package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ArtemShamro/Calendar/events/pkg/api"
)

func main() {
	// SERVICE
	bookingServer := api.NewEventServer()

	bookingServer.Init(context.Background())

	r := http.NewServeMux()

	h := api.HandlerFromMux(bookingServer, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		slog.Debug(err.Error())
		print(err.Error())
	}
}

