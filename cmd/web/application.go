package main

import (
	"log"
	"text/template"

	"github.com/vlkvch/bloggee/internal/models"
)

type application struct {
	infoLog        *log.Logger
	errorLog       *log.Logger
	posts          *models.PostModel
	templateCache  map[string]*template.Template
}
