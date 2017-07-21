package entities

type Schedule []Route

type Route struct {
	Day       string `json:"day"`
	Departure string `json:"departure_at"`
	Arrival   string `json:"arrival_at"`
}
