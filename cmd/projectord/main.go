package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/charmbracelet/log"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	projectorHttp "github.com/OpenSlides/openslides-projector-service/pkg/http"
	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
)

type config struct {
	Bind                 string `env:"BIND" envDefault:":9051"`
	Development          bool   `env:"OPENSLIDES_DEVELOPMENT" envDefault:"false"`
	PostgresHost         string `env:"DATABASE_HOST" envDefault:"localhost"`
	PostgresPort         string `env:"DATABASE_PORT" envDefault:"5432"`
	PostgresDatabase     string `env:"DATABASE_NAME" envDefault:"openslides"`
	PostgresUser         string `env:"DATABASE_USER" envDefault:"openslides"`
	PostgresPasswordFile string `env:"DATABASE_PASSWORD_FILE" envDefault:"/run/secrets/postgres_password"`
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

	ds, err := getDatabase(cfg)
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}

	projectorPool := projector.NewProjectorPool(ds)

	serverMux := http.NewServeMux()
	projectorHandler := projectorHttp.ProjectorHttp{
		ServerMux: serverMux,
		DS:        ds,
		Projector: projectorPool,
	}
	projectorHandler.RegisterRoutes()
	fileHandler := http.StripPrefix("/system/projector/static/", http.FileServer(http.Dir("static")))
	serverMux.Handle("/system/projector/static/", fileHandler)

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

func getDatabase(cfg config) (*datastore.Datastore, error) {
	password, err := parseSecretsFile(cfg.PostgresPasswordFile)
	if err != nil {
		if cfg.Development {
			password = "openslides"
		} else {
			return nil, fmt.Errorf("reading password from secrets: %w", err)
		}
	}

	ds, err := datastore.New(fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(cfg.PostgresUser),
		encodePostgresConfig(password),
		encodePostgresConfig(cfg.PostgresHost),
		encodePostgresConfig(cfg.PostgresPort),
		encodePostgresConfig(cfg.PostgresDatabase),
	))
	if err != nil {
		return nil, fmt.Errorf("creating datastore: %w", err)
	}

	return ds, nil
}

// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

// parseSecretsFile takes a relative path as its argument
// and returns its contents
func parseSecretsFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
