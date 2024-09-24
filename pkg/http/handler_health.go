package http

import (
	"fmt"
	"net/http"
)

func (s *ProjectorHttp) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"healthy": true}`)
	}
}
