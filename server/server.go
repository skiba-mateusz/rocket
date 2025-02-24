package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Server struct {
	port      int
	templates *template.Template
	mu        sync.Mutex
}

func NewServer(port int) *Server {
	return &Server{
		port:      port,
		templates: nil,
	}
}

func (s *Server) Run() error {
	s.loadTemplates()
	http.HandleFunc("/", s.pageHandler)
	fmt.Printf("server started on http://localhost:%d\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func (s *Server) pageHandler(w http.ResponseWriter, r *http.Request) {
	page := strings.Trim(r.URL.Path, "/")
	if page == "" || page == "/" {
		page = "index"
	}
	if err := s.renderPage(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) renderPage(w http.ResponseWriter, page string) error {
	pageFile := fmt.Sprintf("content/%s.md", page)

	content, err := os.ReadFile(pageFile)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"Title":   page,
		"Content": string(content),
	}

	return s.templates.ExecuteTemplate(w, "base", data)
}

func (s *Server) loadTemplates() {
	s.templates = template.Must(template.ParseGlob("layouts/**/*.html"))
}
