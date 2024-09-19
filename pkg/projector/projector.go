package projector

import (
	"fmt"
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

	projector, err := s.DS.Collection(&models.Projector{}).SetIds(1).AsSingle().Run()
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(*projector.(models.Projector).Name)
	}
}
