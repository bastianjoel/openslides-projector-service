package http

import (
	"net/http"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
)

type ProjectorHttp struct {
	ServerMux *http.ServeMux
	DS        *datastore.Datastore
}

func (s *ProjectorHttp) RegisterRoutes() {
	s.ServerMux.HandleFunc("/system/projector/health", s.HealthHandler())
	s.ServerMux.HandleFunc("/system/projector/get/{id}", s.ProjectorGetHandler())
}
