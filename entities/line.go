package entities

import (
	_ "github.com/go-sql-driver/mysql"
)

// Line describes a line entity
type Line struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	FerryID     string   `json:"ferry_id"`
	ADockID     string   `json:"a_dock_id"`
	BDockID     string   `json:"b_dock_id"`
	Schedule    Schedule `json:"schedule"`
}

func getLines(page int, limit int) ([]*Dock, error) {

	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	docks := make([]*Dock, 0)

	rows, err := db.Query("SELECT `id`, `name`, `longitude`, `latitude` FROM `docks` ORDER BY `name` DESC LIMIT ?, ?", off, limit)
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
