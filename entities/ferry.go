package entities

import "github.com/sotoz/Ferrytale/database"

// Ferry describes the ferry entity.
type Ferry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetFerries will fetch the existing ferries from the database.
func GetFerries(page int, limit int) ([]*Ferry, error) {
	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	ferries := make([]*Ferry, 0)

	rows, err := database.DBCon.Query("SELECT `ferries`.`id`, `ferries`.`name` FROM `ferries` ORDER BY `created_at` DESC LIMIT ?, ?", off, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ferry Ferry
		err := rows.Scan(
			&ferry.ID,
			&ferry.Name,
		)
		if err != nil {
			return nil, err
		}
		ferries = append(ferries, &ferry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ferries, nil
}
