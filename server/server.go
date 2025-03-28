package server

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/logger"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	logger logger.Logger
	port   int
}

func NewServer(logger logger.Logger, port int) *Server {
	return &Server{
		mux:    http.NewServeMux(),
		logger: logger,
		port:   port,
	}
}

func (s *Server) Run() error {
	listenAddr := fmt.Sprintf(":%d", s.port)

	s.logger.Info("Server is listening on http://localhost%s", listenAddr)

	s.mux.Handle("/", http.FileServer(http.Dir("./public")))

	return http.ListenAndServe(listenAddr, s.mux)
}
