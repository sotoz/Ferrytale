package entities

import "github.com/sotoz/ferrytale/database"

// Route defines a Route struct with Day,Departure and Arrival time.
type Route struct {
	Day       string `json:"day"`
	Departure string `json:"departure_at"`
	Arrival   string `json:"arrival_at"`
}

// GetRoutes will return an array of Routes.
func GetRoutes(lineID string) ([]*Route, error) {
	routes := make([]*Route, 0)

	rows, err := database.DBCon.Query("SELECT `id`, `line_id`, `day`, `departure_at`, `arrival_at` FROM `schedules` WHERE `schedules`.`line_id`= ?", lineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var route Route
		err := rows.Scan(
			&route.Day,
			&route.Departure,
			&route.Arrival,
		)
		if err != nil {
			return nil, err
		}
		routes = append(routes, &route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}
