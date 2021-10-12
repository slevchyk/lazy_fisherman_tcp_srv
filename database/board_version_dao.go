package database

import (
	"database/sql"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

const (
	selectBoardVersionQuery = `
		SELECT
			bv.id,
			bv.pcb_type,
			bv.cpu,
			bv.description
		FROM 
			board_versions bv`
)

func SelectBoardVersion(db *sql.DB) (*sql.Rows, error) {
	return db.Query(selectBoardVersionQuery)
}

func SelectBoardVersionById(db *sql.DB, id int64) (*sql.Rows, error) {
	return db.Query(selectBoardVersionQuery+`
		WHERE
			bv.id=$1`, id)
}

func ScanBoardVersion(rows *sql.Rows, bv *models.BoardVersion) error {
	return rows.Scan(
		&bv.Id,
		&bv.PcbType,
		&bv.Cpu,
		&bv.Description)
}

func InsertBoardVersion(db *sql.DB, bv models.BoardVersion) (sql.Result, error) {

	if bv.Id == 0 {
		stmt, _ := db.Prepare(`
		INSERT INTO
		board_versions (
				pcb_type,
				cpu,
				description
				)
		VALUES ($1, $2, $3);`)

		return stmt.Exec(bv.PcbType, bv.Cpu, bv.Description)
	}

	stmt, _ := db.Prepare(`
		INSERT INTO
		board_versions (
				id,
				pcb_type,
				cpu,
				description
				)
		VALUES ($1, $2, $3, $4);`)

	return stmt.Exec(bv.Id, bv.PcbType, bv.Cpu, bv.Description)
}

func UpdateBoardVersion(db *sql.DB, bv models.BoardVersion) (sql.Result, error) {

	var err error

	stmt, err := db.Prepare(`
			UPDATE
				board_versions
			SET
				pcb_type = $1,
				cpu = $2,
				description = $3				
			WHERE
				id=$4
			`)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(bv.PcbType, bv.Cpu, bv.Description, bv.Id)

	return res, err
}
