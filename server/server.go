package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/skiba-mateusz/rocket/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	logger logger.Logger
	port   int
	dir    string
}

func NewServer(logger logger.Logger, port int, dir string) (*Server, error) {
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid port: %d (must be 1-65535)", port)
	}

	return &Server{
		logger: logger,
		port:   port,
		dir:    dir,
	}, nil
}

func (s *Server) Run() error {
	listenAddr := fmt.Sprintf(":%d", s.port)
	s.logger.Info("Server is starting on http://localhost%s", listenAddr)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(s.dir)))

	server := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	go func() {
		sig := <-shutdownChan
		s.logger.Info("Received signal: %v. Shutting down...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			s.logger.Error("Error during shutdown: %v", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %v", err)
	}

	s.logger.Info("Server stopped")

	return nil
}
