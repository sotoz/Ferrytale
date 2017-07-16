package controller

import (
	"flag"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sotoz/Ferrytale/entities"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

func listDocks(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewDocksListResponse(entities.Docks)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
func listFerries(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewFerriesListResponse(entities.Ferries)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func Router() http.Handler {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Amsterdam's Ferries webservice."))
	})

	r.Route("/docks", func(r chi.Router) {
		r.With(paginate).Get("/", listDocks)
	})

	r.Route("/ferries", func(r chi.Router) {
		r.With(paginate).Get("/", listFerries)
	})

	return r

}
