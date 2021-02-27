package render

import (
	"bytes"
	"fmt"
	"github.com/senny-matrix/bookings/pkg/config"
	"github.com/senny-matrix/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig)  {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a page given ResponseWriter and a template name
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// Get the template cache from the app config
		tc = app.TemplateCache
	}else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Println("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Errorf("error writing template to browser: %w", err)
	}
}

// CreateTemplateCache creates a template cache as a map of name to a pointer to a template
func CreateTemplateCache() (map[string]*template.Template,error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl.html")
	if err != nil {
		return myCache,err
	}

	for _, page := range pages{
		name := filepath.Base(page)
		//fmt.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")
		if err != nil {
			return myCache,err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")
			if err != nil {
				return myCache,err
			}
		}
		myCache[name] = ts
	}
	return myCache,nil
}