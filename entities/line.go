package entities

import (
	"time"

	"github.com/sotoz/Ferrytale/database"
)

// Line describes a line entity
type Line struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	FerryID     string   `json:"ferry_id"`
	From        string   `json:"from_dock"`
	To          string   `json:"to_dock"`
	Schedule    Schedule `json:"schedule"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func GetLines(page int, limit int) ([]*Line, error) {
	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	lines := make([]*Line, 0)

	rows, err := database.DBCon.Query("SELECT `id`, `description`, `ferry_id`, `created_at`, `updated_at` FROM `lines` ORDER BY `created_at` DESC LIMIT ?, ?", off, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var line Line
		err := rows.Scan(
			&line.ID,
			&line.Description,
			&line.FerryID,
			&line.CreatedAt,
			&line.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		lines = append(lines, &line)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
