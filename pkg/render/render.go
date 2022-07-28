package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/guicattani/go-course/pkg/config"
	"github.com/guicattani/go-course/pkg/models"
)

var ac *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	ac = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	tc := ac.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not in cache")
		return
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, err
}
