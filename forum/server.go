package forum

import (
	"net/http"
	"time"

	"lincoln.boris/forum/pkg/logger"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	log := logger.NewLogger()

	s.httpServer = &http.Server{
		Addr:           ":4000",
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.InfoLog.Println("Server running on")

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return nil
}
