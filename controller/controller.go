package controller

import (
	"context"
	"net/http"
	"time"

	"log"

	"strconv"

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

var pageCtxKey string

type pageOpts struct {
	Page  int
	Limit int
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var page, limit int
		var err error
		pageParam := r.URL.Query().Get("page")
		limitParam := r.URL.Query().Get("limit")
		if pageParam == "" {
			page = 1
		} else {
			page, err = strconv.Atoi(pageParam)
			if err != nil {
				render.Status(r, http.StatusBadRequest)
				log.Print(err)
			}
		}
		if limitParam == "" {
			limit = 1
		} else {
			limit, err = strconv.Atoi(limitParam)
			if err != nil {
				render.Status(r, http.StatusBadRequest)
			}
		}
		ctx := context.WithValue(r.Context(), pageCtxKey, &pageOpts{
			Page:  page,
			Limit: limit,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func listDocks(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewDocksListResponse(entities.Docks)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func listFerries(w http.ResponseWriter, r *http.Request) {
	pgOpts := r.Context().Value(pageCtxKey).(*pageOpts)
	log.Print("Fetching Ferries")

	ferries, err := entities.GetFerries(pgOpts.Page, pgOpts.Limit)
	if err != nil {
		log.Printf("error: %s", err)
	}

	if err := render.RenderList(w, r, NewFerriesListResponse(ferries)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func listLines(w http.ResponseWriter, r *http.Request) {
	pgOpts := r.Context().Value(pageCtxKey).(*pageOpts)
	log.Print("Fetching Lines")

	lines, err := entities.GetLines(pgOpts.Page, pgOpts.Limit)
	if err != nil {
		log.Printf("error: %s", err)
	}
	if err := render.RenderList(w, r, NewLinesListResponse(lines)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Router is the default controller that has the routes for the application.
func Router() http.Handler {
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
	r.Route("/lines", func(r chi.Router) {
		r.With(paginate).Get("/", listLines)
	})

	return r
}
