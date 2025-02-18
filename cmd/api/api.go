package main

import (
	"agenti/internal/app"
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	httpServer := &http.Server{
		Addr:    serverAddress(),
		Handler: app.Routes(), //register the app routes
	}

	go runServer(httpServer) //run the app server in a goroutine

	awaitShutdownSignal() //wait for shutdown signal from the system

	serverGracefulShutdown(httpServer) //shutdown the app server
}

func serverAddress() string {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}

	host, found := os.LookupEnv("HOST")
	if !found {
		host = "0.0.0.0"
	}

	return net.JoinHostPort(host, port)
}

func runServer(httpServer *http.Server) {

	slog.Info("Starting server ...", "address", httpServer.Addr)

	if err := httpServer.ListenAndServe(); !errors.Is(http.ErrServerClosed, err) {
		slog.Error("error listening", "error", err)
	}
}

func awaitShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChan
	slog.Info("Received shutdown signal")
}

func serverGracefulShutdown(server *http.Server) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("error shutting down server", "err", err)
		os.Exit(1)
	}

	slog.Info("Server gracefully stopped")
}
