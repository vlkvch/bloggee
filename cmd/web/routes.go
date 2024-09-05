package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vlkvch/bloggee/ui"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handle("/static/*", fileServer)

	router.Get("/", app.home)
	router.Get("/about", app.about)
	router.Get("/archive", app.archive)

	router.Get("/posts/{id}", app.postView)
	router.Get("/posts/{id}/*", app.postFiles)

	return app.recoverPanic(app.logRequest(middleware.StripSlashes(secureHeaders(router))))
}
