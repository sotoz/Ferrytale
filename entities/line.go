package entities

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sotoz/Ferrytale/database"
)

// Line describes a line entity
type Line struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	FerryID     string   `json:"ferry_id"`
	ADockID     string   `json:"a_dock_id"`
	BDockID     string   `json:"b_dock_id"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Schedule    Schedule `json:"schedule"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetLines(page int, limit int) ([]*Line, error) {

	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	lines := make([]*Line, 0)

	rows, err := database.DBCon.Query("SELECT `id`, `description`, `ferry_id`, `a_dock_id`, `b_dock_id`, `created_at`, `updated_at` FROM `lines` ORDER BY `created_at` DESC LIMIT ?, ?", off, limit)
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
			&line.ADockID,
			&line.BDockID,
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
