package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CodeMaverick-143/Golang_learning/internal/config"
	"github.com/CodeMaverick-143/Golang_learning/internal/http/handlers/health"
	"github.com/CodeMaverick-143/Golang_learning/internal/http/handlers/student"
	"github.com/CodeMaverick-143/Golang_learning/internal/storage"
	"github.com/CodeMaverick-143/Golang_learning/internal/storage/postgres"
	"github.com/CodeMaverick-143/Golang_learning/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	port := os.Getenv("PORT")
	if port != "" {
		cfg.HTTPServer.Address = "0.0.0.0:" + port
	}

	//database setup
	var storage storage.Storage
	var err error

	if cfg.PostgresDSN != "" {
		storage, err = postgres.New(cfg)
		if err != nil {
			log.Fatal(err)
		}
		slog.Info("Connected to Postgres")
	} else {
		storage, err = sqlite.New(cfg)
		if err != nil {
			log.Fatal(err)
		}
		slog.Info("Connected to SQLite")
	}

	slog.Info("Database setup successfully",
		slog.String("env", cfg.Env),
		slog.String("storage_path", cfg.StoragePath),
		slog.String("http_address", cfg.HTTPServer.Address),
		slog.String("version", "1.0.0"),
	)

	// setup router

	router := http.NewServeMux()

	router.Handle("POST /api/students", student.New(storage))
	router.Handle("GET /api/students/{id}", student.GetByID(storage))
	router.Handle("GET /health", health.New())

	// setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	slog.Info("Server started",
		slog.String("env", cfg.Env),
		slog.String("storage_path", cfg.StoragePath),
		slog.String("http_address", cfg.HTTPServer.Address),
	)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
