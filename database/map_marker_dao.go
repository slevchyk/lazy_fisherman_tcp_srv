package database

import (
	"database/sql"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

const (
	selectMapMarkerQuery = `
		SELECT
			mm.id,
			mm.app_user_id,
			mm.ext_id,
			mm.board_id,
			mm.water_id,
			mm.type,
			mm.lat,
			mm.lng,
			mm.title,
			mm.snippet,
			mm.created_at,
			mm.updated_at			
		FROM 
			map_markers mm`
)

func SelectMapMarker(db *sql.DB) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery)
}

func SelectMapMarkerById(db *sql.DB, id string) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.id=$1`, id)
}

func SelectMapMarkerByOwnerId(db *sql.DB, owner_id string) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.app_user_id=$1`, owner_id)
}

func SelectMapMarkerByIdOwnerId(db *sql.DB, id, owner_id string) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.id=$1
			AND mm.app_user_id=$2`, id, owner_id)
}

func SelectMapMarkerByBoardId(db *sql.DB, board_id string) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.board_id=$1`, board_id)
}

func ScanMapMarker(rows *sql.Rows, mm *models.MapMarker) error {
	return rows.Scan(
		&mm.Id,
		&mm.AppUserId,
		&mm.ExtId,
		&mm.BoardId,
		&mm.WaterId,
		&mm.Type,
		&mm.Lat,
		&mm.Lng,
		&mm.Title,
		&mm.Snippet,
		&mm.UpdatedAt,
		&mm.CreatedAt)
}

func UpdateMapMarker(db *sql.DB, mm models.MapMarker) (sql.Result, error) {

	var err error

	stmt, err := db.Prepare(`
			UPDATE
				map_markers
			SET
				app_user_id = $1,
				ext_id = $2,
				board_id = $3,				
				water_id = $4,				
				type = $5,				
				lat = $6,				
				lng = $7,				
				title = $8,				
				snippet = $9,								
				updated_at = $10
			WHERE
				id=$12
			`)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(mm.AppUserId, mm.ExtId, mm.BoardId, mm.WaterId, mm.Type, mm.Lat, mm.Lng, mm.Title, mm.Snippet, mm.UpdatedAt, mm.Id)

	return res, err
}

func InsertMapMarker(db *sql.DB, mm models.MapMarker) (sql.Result, error) {

	stmt, _ := db.Prepare(`
		INSERT INTO
		map_markers (
				id,
				app_user_id,
				ext_id,
				board_id,
				water_id,
				type,
				lat,
				lng,
				title,
				snippet,
				created_at,
				updated_at
				)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`)

	return stmt.Exec(mm.Id, mm.AppUserId, mm.ExtId, mm.BoardId, mm.WaterId, mm.Type, mm.Lat, mm.Lng, mm.Title, mm.Snippet, mm.CreatedAt, mm.UpdatedAt)
}
