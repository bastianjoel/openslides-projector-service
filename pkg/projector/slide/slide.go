package slide

import (
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projectionRequest struct {
	Projection    *models.Projection
	DB            *datastore.Datastore
	UpdateTrigger <-chan struct{}
}

type SlideRouter struct {
	db     *datastore.Datastore
	Routes map[string]func(*projectionRequest) <-chan string
}

func New(db *datastore.Datastore, projectorID int) *SlideRouter {
	routes := make(map[string]func(*projectionRequest) <-chan string)

	return &SlideRouter{
		Routes: routes,
	}
}

func (r *SlideRouter) SubscribeContent(update <-chan struct{}, projection *models.Projection) <-chan string {
	projectionType := getProjectionType(projection)
	if handler, ok := r.Routes[projectionType]; ok {
		return handler(&projectionRequest{
			Projection:    projection,
			DB:            r.db,
			UpdateTrigger: update,
		})
	}

	return nil
}

func getProjectionType(projection *models.Projection) string {
	if projection.Type != nil {
		return *projection.Type
	}

	collection, _, found := strings.Cut(projection.ContentObjectID, "/")
	if found {
		return collection
	}

	return "unknown"
}
