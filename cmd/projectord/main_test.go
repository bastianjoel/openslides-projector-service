package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-projector-service/pkg/projector"
)

func TestHandlerHealth(t *testing.T) {
	serverMux := http.NewServeMux()
	projectorHandler := projector.Projector{
		ServerMux: serverMux,
	}
	projectorHandler.RegisterRoutes()
	rec := httptest.NewRecorder()

	projectorHandler.ServerMux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health", strings.NewReader(``)))

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "{\"healthy\": true}\n" {
		t.Errorf("expected body %q, got %q", `{"healthy": true}`, rec.Body.String())
	}
}
