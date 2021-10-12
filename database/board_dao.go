package database

import (
	"database/sql"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

const (
	selectBoardQuery = `
		SELECT
			b.id,
			b.board_version_id,
			b.name,
			b.current_firware,
			b.latest_firmware,
			b.app_user_id,
			b.serial_number
		FROM 
			boards b`
)

func SelectBoard(db *sql.DB) (*sql.Rows, error) {
	return db.Query(selectBoardQuery)
}

func SelectBoardById(db *sql.DB, id int64) (*sql.Rows, error) {
	return db.Query(selectBoardQuery+`
		WHERE
			b.id=$1`, id)
}

func SelectBoardByOwnerId(db *sql.DB, app_user_id int64) (*sql.Rows, error) {
	return db.Query(selectBoardQuery+`
		WHERE
			b.app_user_id=$1`, app_user_id)
}

func SelectBoardBySerialNumber(db *sql.DB, serial_number string) (*sql.Rows, error) {
	return db.Query(selectBoardQuery+`
		WHERE
			b.serial_number=$1`, serial_number)
}

func ScanBoard(rows *sql.Rows, b *models.Board) error {
	return rows.Scan(
		&b.Id,
		&b.BoardVersionId,
		&b.Name,
		&b.CurrentFirmware,
		&b.LatestFirmware,
		&b.AppUserId,
		&b.SerialNumber)
}

func InsertBoard(db *sql.DB, b models.Board) (int64, error) {

	var lastInsertId int64
	var err error

	if b.Id == 0 {
		err = db.QueryRow(`
		INSERT INTO
		boards (
				board_version_id,
				name,
				current_firware,
				latest_firmware,
				app_user_id
				)
		VALUES ($1, $2, $3, $4, $5)  RETURNING id`, b.BoardVersionId, b.Name, b.CurrentFirmware, b.LatestFirmware, b.AppUserId).Scan(&lastInsertId)

		return lastInsertId, err
	}

	err = db.QueryRow(`
		INSERT INTO
		boards (
				id,
				board_version_id,
				name,
				current_firware,
				latest_firmware,
				app_user_id
				)
		VALUES ($1, $2, $3, $4, $5, $6)  RETURNING id`, b.Id, b.BoardVersionId, b.Name, b.CurrentFirmware, b.LatestFirmware, b.AppUserId).Scan(&lastInsertId)

	return lastInsertId, err
}

func UpdateBoard(db *sql.DB, b models.Board) (sql.Result, error) {

	var err error

	stmt, err := db.Prepare(`
			UPDATE
				boards
			SET
				board_version_id = $1,
				name = $2,
				current_firware = $3,				
				latest_firmware = $4,
				app_user_id = $5
			WHERE
				id=$6
			`)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(b.BoardVersionId, b.Name, b.CurrentFirmware, b.LatestFirmware, b.AppUserId, b.Id)

	return res, err
}
