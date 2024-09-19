package projector

import (
	"net/http"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
)

type Projector struct {
	ServerMux *http.ServeMux
	DS        *datastore.Datastore
}

func (s *Projector) RegisterRoutes() {
	s.ServerMux.HandleFunc("/health", s.HealthHandler())
	s.ServerMux.HandleFunc("/get/{id}", s.ProjectorGetHandler())
}
