package web

import (
	"context"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
)

// A Handler is a type that handles an http request within our own little mini framework
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures out context
// object for each of out http handlers.
// THIS IS FOR MODIFICATION with any logic and/or data
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// NerApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}
}

// Handle sets a handler fucntion for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
	//middleware specific to the route
	handler = wrapMiddleware(mw, handler)
	//middleware spicific app level middleware around previous
	handler = wrapMiddleware(a.mw, handler)
	h := func(w http.ResponseWriter, r *http.Request) {

		//ADD any logic here

		if err := handler(r.Context(), w, r); err != nil {
			//TODO
			return
		}

		//ADD any logic here
	}

	a.ContextMux.Handle(method, path, h)
}
