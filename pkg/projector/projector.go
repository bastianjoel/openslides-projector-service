package projector

import (
	"log"
	"net/http"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type Projector struct {
	ServerMux *http.ServeMux
	DS        *datastore.Datastore
}

func (s *Projector) RegisterRoutes() {
	s.ServerMux.HandleFunc("/health", s.HealthHandler())

	_, err := s.DS.Collection(&models.Projector{}).SetIds(1, 2).Run()
	if err != nil {
		log.Print(err)
	}
}
