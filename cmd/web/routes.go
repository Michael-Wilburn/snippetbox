package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux
func (app *application) routes() http.Handler {
	// Initialize the router.
	router := httprouter.New()

	// Create a handler fucntion which wraps our notFound() helper, and then
	// assing is as the custom handler for 404 Not Found responses. You can also
	// sey a custom handler for 405 Method Not Allowed responses by setting
	// router.MethodNotAllow in the same way too.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Update the pattern for the route for the static files.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Create the routes using the appropiate methods, patterns and handlers.
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Returnd the 'standard' middlew chain followed by the servemux.
	return standard.Then(router)
}
