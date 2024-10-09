package slide

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

type projectionRequest struct {
	Projection *models.Projection
	DB         *datastore.Datastore
}

type projectionUpdate struct {
	Content string
}

type SlideRouter struct {
	ctx    context.Context
	db     *datastore.Datastore
	Routes map[string]func(*projectionRequest) <-chan string
}

func New(ctx context.Context, db *datastore.Datastore) *SlideRouter {
	routes := make(map[string]func(*projectionRequest) <-chan string)

	return &SlideRouter{
		ctx:    ctx,
		db:     db,
		Routes: routes,
	}
}

func (r *SlideRouter) SubscribeContent(addProjection <-chan int, removeProjection <-chan int) <-chan *projectionUpdate {
	updateChannel := make(chan *projectionUpdate)
	closeChannels := make(map[int]chan struct{})

	go func() {
		for {
			select {
			case <-r.ctx.Done():
				close(updateChannel)
				for id := range closeChannels {
					closeChannels[id] <- struct{}{}
					close(closeChannels[id])
				}

				return
			case id := <-addProjection:
				if _, ok := closeChannels[id]; !ok {
					closeChannels[id] = make(chan struct{})
					go r.subscribeProjection(id, updateChannel, closeChannels[id])
				}
			case id := <-removeProjection:
				closeChannels[id] <- struct{}{}
				close(closeChannels[id])
				delete(closeChannels, id)
			}
		}
	}()

	return updateChannel
}

func (r *SlideRouter) subscribeProjection(id int, updateChannel chan<- *projectionUpdate, closeSubscription <-chan struct{}) {
	projection, err := datastore.Collection(r.db, &models.Projection{}).SetIds(id).SetFields("id", "content_object_id", "type").GetOne()
	if err != nil {
		fmt.Println("getting projection type and content object from db: %w", err)
		return
	}

	projectionType := getProjectionType(projection)
	if handler, ok := r.Routes[projectionType]; ok {
		projectionChan := handler(&projectionRequest{
			Projection: projection,
			DB:         r.db,
		})

		for {
			select {
			case <-closeSubscription:
				return
			case projectionContent := <-projectionChan:
				updateChannel <- &projectionUpdate{
					Content: projectionContent,
				}
			}
		}
	} else {
		fmt.Println(fmt.Errorf("unknown projection type %s", projectionType))
	}
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
