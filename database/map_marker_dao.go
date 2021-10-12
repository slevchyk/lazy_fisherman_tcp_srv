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

func SelectMapMarkerByAppUserId(db *sql.DB, appUserId int64) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.app_user_id=$1`, appUserId)
}

func SelectMapMarkerByIdAppUserId(db *sql.DB, id, appUserId int64) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.id=$1
			AND mm.app_user_id=$2`, id, appUserId)
}

func SelectMapMarkerByBoardId(db *sql.DB, boardId int64) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.board_id=$1`, boardId)
}

func SelectMapMarkerByBoardIdWaterId(db *sql.DB, boardId, waterId int64) (*sql.Rows, error) {
	return db.Query(selectMapMarkerQuery+`
		WHERE
			mm.board_id=$1
			AND ww.water_id=$2`, boardId, waterId)
}

func ScanMapMarker(rows *sql.Rows, mm *models.MapMarker) error {
	return rows.Scan(
		&mm.Id,
		&mm.AppUserId,
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
				board_id = $2,				
				water_id = $3,				
				type = $4,				
				lat = $5,				
				lng = $6,				
				title = $7,				
				snippet = $8,								
				updated_at = $9
			WHERE
				id=$12
			`)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(mm.AppUserId, mm.BoardId, mm.WaterId, mm.Type, mm.Lat, mm.Lng, mm.Title, mm.Snippet, mm.UpdatedAt, mm.Id)

	return res, err
}

func InsertMapMarker(db *sql.DB, mm models.MapMarker) (int64, error) {

	var lastInsertId int64
	var err error

	if mm.Id == 0 {
		err = db.QueryRow(`
		INSERT INTO
		map_markers (
				app_user_id,				
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`, mm.AppUserId, mm.BoardId, mm.WaterId, mm.Type, mm.Lat, mm.Lng, mm.Title, mm.Snippet, mm.CreatedAt, mm.UpdatedAt).Scan(&lastInsertId)

		return lastInsertId, err
	}

	err = db.QueryRow(`
		INSERT INTO
		map_markers (
				id,
				app_user_id,				
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)  RETURNING id`, mm.Id, mm.AppUserId, mm.BoardId, mm.WaterId, mm.Type, mm.Lat, mm.Lng, mm.Title, mm.Snippet, mm.CreatedAt, mm.UpdatedAt).Scan(&lastInsertId)

	return lastInsertId, err
}
