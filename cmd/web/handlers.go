package main

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/vlkvch/bloggee/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	latestPosts := []*models.Post{}

	for i := range 3 {
		latestPosts = append(latestPosts, posts[i])
	}

	app.render(w, http.StatusOK, "home", &templateData{
		Posts: latestPosts,
	})
}

func (app *application) archive(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "archive", &templateData{
		Posts: posts,
	})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := new(templateData)

	app.render(w, http.StatusOK, "about", data)
}

func (app *application) postView(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := app.posts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoPost) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	app.render(w, http.StatusOK, "view", &templateData{
		Post: post,
	})
}

func (app *application) postFiles(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	postDir := filepath.Join(fmt.Sprint(app.posts.Dir), id)

	fileServer := http.FileServer(http.Dir(postDir))
	http.StripPrefix(fmt.Sprintf("/posts/%s", id), fileServer).ServeHTTP(w, r)
}
