package projector

import (
	"fmt"
	"net/http"
)

func (s *Projector) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"healthy": true}`)
	}
}
