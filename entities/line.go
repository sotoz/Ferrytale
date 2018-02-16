package entities

import (
	"time"

	"github.com/sotoz/ferrytale/database"
)

// Line describes a line entity
type Line struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Ferry       string     `json:"ferry"`
	From        string     `json:"from_dock"`
	To          string     `json:"to_dock"`
	Routes      []Route    `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

// GetLines will fetch the existing ferry lines from the database.
func GetLines(page int, limit int) ([]*Line, error) {
	var off int
	if page < 2 {
		off = 0
	} else {
		off = (page - 1) * limit
	}

	lines := make([]*Line, 0)

	rows, err := database.DBCon.Query("SELECT `lines`.`id`, `lines`.`description`, (SELECT `docks`.`name` FROM `docks` WHERE `docks`.`id`=`lines`.`a_dock_id`) AS docka,(SELECT `docks`.`name` FROM `docks` WHERE `docks`.`id`=`lines`.`b_dock_id`) AS dockb, (SELECT `ferries`.`name` FROM `ferries` WHERE `ferries`.`id`=`lines`.`ferry_id`) AS ferry FROM `lines` ORDER BY `created_at` DESC LIMIT ?, ?", off, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var line Line
		err := rows.Scan(
			&line.ID,
			&line.Description,
			&line.From,
			&line.To,
			&line.Ferry,
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

// GetLine returns the whole line record from the database.
func GetLine(lineID string) (*Line, error) {
	var line Line

	row := database.DBCon.QueryRow("SELECT `lines`.`id`, `lines`.`description`, (SELECT `docks`.`name` FROM `docks` WHERE `docks`.`id`=`lines`.`a_dock_id`) AS docka,(SELECT `docks`.`name` FROM `docks` WHERE `docks`.`id`=`lines`.`b_dock_id`) AS dockb, (SELECT `ferries`.`name` FROM `ferries` WHERE `ferries`.`id`=`lines`.`ferry_id`) AS ferry FROM `lines` WHERE `lines`.`id`=? LIMIT 1", lineID)
	if row == nil {
		return &line, nil
	}

	err := row.Scan(
		&line.ID,
		&line.Description,
		&line.From,
		&line.To,
		&line.Ferry,
	)

	return &line, err
}
