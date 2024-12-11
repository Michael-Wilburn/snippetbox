package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/Michael-Wilburn/snippetbox/internal/models"
	"github.com/Michael-Wilburn/snippetbox/ui"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool   // Add an IsAuthentiated field to the templateData struct.
	CSRFToken       string // Add a CSRFToken field.
}

// Create a humanDate function which returns a nicely formated string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a templete.FuncMap object and store it in a global variable. This
// is essentially a string-keyd map which acts as a lookup between the names of
// our custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl". This will essetially gives
	// us a slice of all the filepath for our application 'page' templates
	// like: [ui/html/pages/home.tmpl ui/html/view.tmpl]
	/*pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}*/

	// Use fs.Glob() to get a slice of all filepaths in ui.Files embedded
	// filesystem which match the pattern 'html/pages/*.tmpl'. This essentially
	// give us a slice of all the 'page' templates for the application, just
	// like before.
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.tmpl) from the full filepath
		// and assign it to the name varible.
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		/* ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}*/

		// Create a slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		// Use ParseFS() instead of ParseFiles() to parse the template files
		// from the ui.Files embedded filesystem.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		// ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		// if err != nil {
		// 	return nil, err
		// }

		// Call ParseFiles() *on this template set* to add the page template.
		// ts, err = ts.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}

	// Return the map.
	return cache, nil

}
