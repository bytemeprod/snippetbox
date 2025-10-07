package main

import (
	"flag"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/bytemeprod/snippetbox/internal/app"
	"github.com/bytemeprod/snippetbox/pkg/prettylog"
)

// Declaring our config
var cfg app.Config

// Parsing flags to config
func init() {
	flag.StringVar(&cfg.Addr, "addr", ":4040", "HTTP network address to start server")
	flag.StringVar(&cfg.StaticDir, "staticDir", "./frontend/static", "Path to static files")
	flag.Parse()
}

func main() {
	plog := setupPrettyLogger(os.Stdout)

	plog.Debug("Debug messages enabled")

	app := app.NewApplication(plog, cfg)

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: app.SetupRoutes(),
	}

	plog.Info("Server starting...", slog.String("port", cfg.Addr))

	if err := server.ListenAndServe(); err != nil {
		plog.Error("Error", slog.String("error", err.Error()))
	}
}

func setupPrettyLogger(w io.Writer) *slog.Logger {
	levelEnv := os.Getenv("LOG_LEVEL")

	var level slog.Level

	switch levelEnv {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	h := prettylog.NewPrettyHandler(w, slog.HandlerOptions{
		Level: level,
	})
	return slog.New(h)
}
