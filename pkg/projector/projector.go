package projector

import (
	"net/http"
)

type Projector struct {
	ServerMux *http.ServeMux
}

func (s *Projector) RegisterRoutes() {
	s.ServerMux.HandleFunc("/health", s.HealthHandler())
}
