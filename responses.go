package main

import (
	"net/http"

	"github.com/go-chi/render"
)

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type DockListResponse []*DockResponse
type DockResponse struct {
	*Dock
}

func NewDocksListResponse(docks []*Dock) []render.Renderer {
	list := []render.Renderer{}
	for _, dock := range docks {
		list = append(list, NewDockResponse(dock))
	}
	return list
}

func NewDockResponse(dock *Dock) *DockResponse {
	resp := &DockResponse{Dock: dock}

	return resp
}

func (rd *DockResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type FerryListResponse []*FerryResponse
type FerryResponse struct {
	*Ferry
}

func NewFerriesListResponse(ferries []*Ferry) []render.Renderer {
	list := []render.Renderer{}
	for _, ferry := range ferries {
		list = append(list, NewFerryResponse(ferry))
	}
	return list
}

func NewFerryResponse(ferry *Ferry) *FerryResponse {
	resp := &FerryResponse{Ferry: ferry}

	return resp
}

func (rd *FerryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
