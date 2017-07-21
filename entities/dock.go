package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sotoz/Ferrytale/database"
)

// Dock describes a dock entity
type Dock struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// Docks list has all the fixture data for docks. @todo these will be taken from the database
var Docks = []*Dock{
	{ID: "1", Name: "Westerdoksdijk", Latitude: "Hi", Longitude: "hi"},
	{ID: "2", Name: "Amsterdam Central", Latitude: "222", Longitude: "333"},
}

func getDocks(page int, limit int) ([]*Dock, error) {
	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	docks := make([]*Dock, 0)

	rows, err := database.DBCon.Query("SELECT `id`, `name`, `longitude`, `latitude` FROM `docks` ORDER BY `name` DESC LIMIT ?, ?", off, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var dock Dock
		err := rows.Scan(
			&dock.ID,
			&dock.Name,
			&dock.Longitude,
			&dock.Latitude,
		)
		if err != nil {
			return nil, err
		}
		docks = append(docks, &dock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return docks, nil

}
