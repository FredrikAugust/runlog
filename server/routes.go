package server

import (
	"net/http"

	"github.com/fredrikaugust/runlog/routes"
)

func (s *Server) SetupRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", routes.Health)
	mux.HandleFunc("/register", routes.Register(s.db))
	mux.HandleFunc("/run/upload", routes.Upload(s.db))

	s.httpServer.Handler = mux
}
