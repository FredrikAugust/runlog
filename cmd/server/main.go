package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fredrikaugust/runlog/server"
	"github.com/fredrikaugust/runlog/storage"
	"golang.org/x/sync/errgroup"
)

func main() {
	slog.Info("starting program")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)

	db, err := storage.NewDB(ctx, &storage.DBOptions{
		ConnectionURL: "postgres://postgres:localdev@localhost:5432/postgres?sslmode=disable",
	})
	if err != nil {
		slog.Error("could not connect to db", "error", err.Error())
	}

	server := server.NewServer(&server.ServerOptions{
		Address: "0.0.0.0:8080",
		DB:      db,
	})
	server.SetupRoutes()

	eg.Go(func() error {
		return server.Start(ctx)
	})

	err = eg.Wait()
	if err != nil {
		slog.Error("error occurred in group", "error", err.Error())
	}

	slog.Info("shutdown complete")
}
