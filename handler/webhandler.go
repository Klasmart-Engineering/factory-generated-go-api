package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go-template/models"
	"net/http"
)

func healthcheck(router chi.Router) {
	router.Get("/", dbHealthcheck)
}

func liveness(router chi.Router) {
	router.Get("/", applicationLiveness)
}

func example(router chi.Router) {
	router.Get("/", exampleGet)
}

func applicationLiveness(w http.ResponseWriter, r *http.Request) {

	health := models.Healthcheck{Up: true, Details: "Application up"}
	if err := render.Render(w, r, &health); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func exampleGet(w http.ResponseWriter, r *http.Request) {

	exampleStructs, err := dbInstance.ExampleGet()

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, exampleStructs); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func dbHealthcheck(w http.ResponseWriter, r *http.Request) {
	up, err := dbInstance.Healthcheck()

	if err != nil || !up {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	health := models.Healthcheck{Up: true, Details: "Successfully connected to the database"}
	if err := render.Render(w, r, &health); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
