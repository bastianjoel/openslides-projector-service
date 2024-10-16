package slide

import "context"

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	go func() {
		projection := req.Projection
		content <- "Topic: " + projection.ContentObjectID
	}()

	return content, nil
}
