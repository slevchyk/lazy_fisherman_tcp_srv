package database

import (
	"database/sql"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

func SelectAppUserByPhonePassword(db *sql.DB, phone, password string) (*sql.Rows, error) {
	return db.Query(`
		SELECT
			au.id,
			au.name,
			au.password,
			au.phone,
			au.email,
			au.token,
			au.is_blocked,
			au.is_administrator
		FROM 
			app_users au
		WHERE
			au.phone=$1
			and au.password>=$2`, phone, password)
}

func SelectAppUserByPhone(db *sql.DB, phone string) (*sql.Rows, error) {
	return db.Query(`
		SELECT
			au.id,
			au.name,
			au.password,
			au.phone,
			au.email,
			au.token,
			au.is_blocked,
			au.is_administrator
		FROM 
			app_users au
		WHERE
			au.phone=$1`, phone)
}

func SelectAppUserById(db *sql.DB, id int64) (*sql.Rows, error) {
	return db.Query(`
		SELECT
			au.id,
			au.name,
			au.password,
			au.phone,
			au.email,
			au.token,
			au.is_blocked,
			au.is_administrator
		FROM 
			app_users au
		WHERE
			au.id=$1`, id)
}

func SelectAppUserByToken(db *sql.DB, token string) (*sql.Rows, error) {
	return db.Query(`
		SELECT
			au.id,
			au.name,
			au.password,
			au.phone,
			au.email,
			au.token,
			au.is_blocked,
			au.is_administrator
		FROM 
			app_users au
		WHERE
			au.token=$1`, token)
}

func SelectAppUser(db *sql.DB) (*sql.Rows, error) {
	return db.Query(`
		SELECT
			au.id,
			au.name,
			au.password,
			au.phone,
			au.email,
			au.token,
			au.is_blocked,
			au.is_administrator
		FROM 
			app_users au
		WHERE`)
}

func ScanAppUser(rows *sql.Rows, au *models.AppUser) error {
	return rows.Scan(
		&au.Id,
		&au.Name,
		&au.Password,
		&au.Phone,
		&au.Email,
		&au.Token,
		&au.IsBlocked,
		&au.IsAdministrator)
}

func InsertAppUser(db *sql.DB, au models.AppUser) (sql.Result, error) {

	if au.Id == 0 {
		stmt, _ := db.Prepare(`
		INSERT INTO
			app_users (				
				name,
				password,
				phone,
				email,
				token,
				is_blocked,
				is_administrator
				)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`)

		return stmt.Exec(au.Name, au.Password, au.Phone, au.Email, au.Token, au.IsBlocked, au.IsAdministrator)
	}

	stmt, _ := db.Prepare(`
		INSERT INTO
			app_users (
				id,
				name,
				password,
				phone,
				email,
				token,
				is_blocked,
				is_administrator
				)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`)

	return stmt.Exec(au.Id, au.Name, au.Password, au.Phone, au.Email, au.Token, au.IsBlocked, au.IsAdministrator)
}

func UpdateAppUser(db *sql.DB, au models.AppUser) (sql.Result, error) {

	var err error

	stmt, err := db.Prepare(`
			UPDATE
				app_users
			SET
				name = $1,
				password = $2,
				phone = $3,
				email = $4,
				token = $5,
				is_blocked = $6,
				is_administrator = $7
			WHERE
				id=$8
			`)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(au.Name, au.Password, au.Phone, au.Email, au.Token, au.IsBlocked, au.IsAdministrator, au.Id)

	return res, err
}
