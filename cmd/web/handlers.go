package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Michael-Wilburn/snippetbox/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Initialize a map to hold any validation erros for the form fields.
	fieldErrors := make(map[string]string)

	// Check that the title value is not blank and is not more than 100
	// characters long. If it fails either of those checks, add a message to error map using the field name as the key.
	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot be blank."
	} else if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field cannot be more than 100 characters long."
	}

	// Check that the content value isn't blank.
	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field cannot be blank."
	}

	// Check the expires values matches one of the permitted values (1,7,365)
	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7, or 365"
	}

	// If there are any errors, dump in a plain text HTTP response
	// and return from the handler.
	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
