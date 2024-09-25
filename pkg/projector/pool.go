package projector

import (
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-projector-service/pkg/datastore"
)

type ProjectorPool struct {
	mu         sync.Mutex
	projectors map[int]*projector
	db         *datastore.Datastore
}

func NewProjectorPool(db *datastore.Datastore) *ProjectorPool {
	return &ProjectorPool{
		db:         db,
		projectors: make(map[int]*projector),
	}
}

func (pool *ProjectorPool) readOrCreateProjector(id int) (*projector, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if projector, ok := pool.projectors[id]; ok {
		return projector, nil
	}

	projector, err := newProjector(id, pool.db)
	if err != nil {
		return nil, fmt.Errorf("error creating new projector: %w", err)
	}

	pool.projectors[id] = projector
	return projector, nil
}

func (pool *ProjectorPool) GetProjectorContent(id int) (*string, error) {
	projector, err := pool.readOrCreateProjector(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector content: %w", err)
	}

	return &projector.Content, nil
}

func (pool *ProjectorPool) SubscribeProjectorContent(id int) (<-chan projectorUpdateEvent, error) {
	projector, err := pool.readOrCreateProjector(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving projector channel: %w", err)
	}

	return projector.Channel, nil
}
