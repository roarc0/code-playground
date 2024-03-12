package assets

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/rs/zerolog/log"
)

//go:embed static/*
var staticFS embed.FS

// Static return the static files to be served directly by the http router
func Static() (fs.FS, error) {
	return fs.Sub(staticFS, "static")
}

//go:embed templates/*
var templatesFS embed.FS

// Templates returns the html templates to be rendered
func Templates() *template.Template {
	templ, err := template.New("").ParseFS(templatesFS, "templates/*.tmpl")
	if err != nil {
		log.Fatal().Err(err).Msg("could not retrieve templates")
	}
	return templ
}
