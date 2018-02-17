package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/sotoz/ferrytale/entities"
)

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type RouteListResponse []*RouteResponse
type RouteResponse struct {
	*entities.Route
}

func NewRoutesListResponse(routes []*entities.Route) []render.Renderer {
	list := []render.Renderer{}
	for _, route := range routes {
		list = append(list, NewRouteResponse(route))
	}
	return list
}

func NewRouteResponse(route *entities.Route) *RouteResponse {
	resp := &RouteResponse{Route: route}

	return resp
}

func (rd *RouteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshaled and sent across the wire
	return nil
}

type NextDepartureResponse struct {
	NextDepartureTime      time.Time `json:"next_departure_time"`
	MinutesBeforeDeparture string    `json:"minutes_before_departure"`
}

func NewNextDepartureResponse(nextdep *time.Time, d time.Duration) *NextDepartureResponse {
	resp := &NextDepartureResponse{NextDepartureTime: *nextdep, MinutesBeforeDeparture: d.String()}

	return resp
}

func (rd *NextDepartureResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshaled and sent across the wire
	return nil
}

type DockListResponse []*DockResponse
type DockResponse struct {
	*entities.Dock
}

func NewDocksListResponse(docks []*entities.Dock) []render.Renderer {
	list := []render.Renderer{}
	for _, dock := range docks {
		list = append(list, NewDockResponse(dock))
	}
	return list
}

func NewDockResponse(dock *entities.Dock) *DockResponse {
	resp := &DockResponse{Dock: dock}

	return resp
}

func (rd *DockResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshaled and sent across the wire
	return nil
}

type FerryListResponse []*FerryResponse
type FerryResponse struct {
	*entities.Ferry
}

func NewFerriesListResponse(ferries []*entities.Ferry) []render.Renderer {
	list := []render.Renderer{}
	for _, ferry := range ferries {
		list = append(list, NewFerryResponse(ferry))
	}
	return list
}

func NewFerryResponse(ferry *entities.Ferry) *FerryResponse {
	resp := &FerryResponse{Ferry: ferry}

	return resp
}

func (rd *FerryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type LineListResponse []*LineResponse
type LineResponse struct {
	*entities.Line
}

func NewLinesListResponse(lines []*entities.Line) []render.Renderer {
	list := []render.Renderer{}
	for _, line := range lines {
		list = append(list, NewLineResponse(line))
	}
	return list
}

func NewLineResponse(line *entities.Line) *LineResponse {
	resp := &LineResponse{Line: line}

	return resp
}

func (rd *LineResponse) Render(w http.ResponseWriter, r *http.Request) error {
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

type LineAPIResponse struct {
	Data []*entities.Line
}
