package main

import "snippetbox.jrvldam.github.com/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
