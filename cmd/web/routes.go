package main

import (
	"net/http"

	"github.com/vlkvch/bloggee/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("GET /static/*", fileServer)

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /archive", app.archive)

	mux.HandleFunc("GET /posts/{id}", app.postView)
	mux.HandleFunc("GET /posts/{id}/*", app.postFiles)

	return app.recoverPanic(app.logRequest(stripSlashes(secureHeaders(mux))))
}
