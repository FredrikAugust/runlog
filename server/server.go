package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/fredrikaugust/runlog/storage"
)

type ServerOptions struct {
	Address string
	DB      *storage.DB
}

type Server struct {
	httpServer *http.Server
	db         *storage.DB
}

func NewServer(so *ServerOptions) *Server {
	srv := &http.Server{
		Addr: so.Address,
	}

	return &Server{
		httpServer: srv,
		db:         so.DB,
	}
}

func (s *Server) Stop(ctx context.Context) error {
	slog.Info("shutting down http server")
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Start(ctx context.Context) error {
	slog.Info("starting server on :8080")

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Stop(shutdownCtx); err != nil {
			slog.Error("failed to shut down http server", "error", err.Error())
		}
	}()

	err := s.httpServer.ListenAndServe()
	if err == http.ErrServerClosed {
		// This is expected
		return nil
	}
	return err
}
