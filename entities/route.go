package entities

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sotoz/ferrytale/database"
)

// Route defines a Route struct with Day,Departure and Arrival time.
type Route struct {
	Day           string         `json:"day"`
	Departure     mysql.NullTime `json:"departure_at"`
	Arrival       mysql.NullTime `json:"arrival_at"`
	NextDeparture time.Duration
}

// GetRoutes will return an array of Routes.
func GetRoutes(lineID string) ([]*Route, error) {
	routes := make([]*Route, 0)

	rows, err := database.DBCon.Query("SELECT `day`, `departure_at`, `arrival_at` FROM `schedules` WHERE `schedules`.`line_id`= ? ORDER BY `departure_at` ASC;", lineID)
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
