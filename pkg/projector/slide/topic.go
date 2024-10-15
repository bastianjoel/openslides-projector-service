package slide

import "context"

func TopicSlideHandler(ctx context.Context, req *projectionRequest) (<-chan string, error) {
	content := make(chan string)
	go func() {
		content <- "Topic"
	}()

	return content, nil
}
