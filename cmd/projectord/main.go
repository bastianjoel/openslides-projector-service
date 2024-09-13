package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/charmbracelet/log"

	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
)

type config struct {
	Bind  string `env:"BIND" envDefault:"127.0.0.1:8080"`
	Debug bool   `env:"DEBUG" envDefault:"false"`
}

func main() {
	log.Info("Starting...")

	if err := run(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	log.Info("Stopped")
}

func run() error {
	log.Info("Parsing environment variables...")
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	serverMux := http.NewServeMux()
	projectorHandler := projector.Projector{
		ServerMux: serverMux,
	}
	projectorHandler.RegisterRoutes()

	log.Infof("Starting server on %s", cfg.Bind)
	srv := &http.Server{
		Addr:    cfg.Bind,
		Handler: serverMux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return nil
}
