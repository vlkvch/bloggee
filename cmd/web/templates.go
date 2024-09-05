package main

import (
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/vlkvch/bloggee/internal/models"
	"github.com/vlkvch/bloggee/ui"
)

type templateData struct {
	Post  *models.Post
	Posts []*models.Post
	Form  any
	Flash string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		baseName := filepath.Base(page)
		name := strings.TrimSuffix(baseName, filepath.Ext(baseName))

		ts, err := template.ParseFS(ui.Files, "html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(ui.Files, "html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(ui.Files, page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
