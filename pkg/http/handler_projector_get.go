package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/OpenSlides/openslides-projector-service/pkg/models"
)

func (s *ProjectorHttp) ProjectorGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Fprintln(w, `{"healthy": true}`)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(`{"error": true, "msg": "Projector id invalid"}`)); err != nil {
				log.Print(err)
			}
			return
		}

		projector, err := s.DS.Collection(&models.Projector{}).SetIds(id).AsSingle().Run()
		if err != nil {
			log.Print(err)
			return
		}

		tmpl, err := template.ParseFiles("templates/projector.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Projector": projector,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
