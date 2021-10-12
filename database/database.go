package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/slevchyk/lazy_fisherman_tcp_srv/models"
)

func ConnectDb(cfg models.Config) (*sql.DB, error) {

	dbConnection := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Server, cfg.DB.Name)
	db, err := sql.Open("postgres", dbConnection)

	return db, err
}

func InitDb(db *sql.DB) {

	var err error

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS app_users (
			id TEXT PRIMARY KEY,
			name TEXT DEFAULT '',			
			password TEXT DEFAULT '',						
			phone TEXT DEFAULT '',
			email TEXT DEFAULT '',
			token TEXT DEFAULT '',			
			is_blocked BOOLEAN DEFAULT false,
			is_administrator BOOLEAN DEFAULT false );`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS board_versions (
			id TEXT PRIMARY KEY,			
			pcb_type TEXT DEFAULT '',						
			cpu TEXT DEFAULT '',
			description TEXT DEFAULT '' );`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS boards (
			id TEXT PRIMARY KEY,			
			board_version_id TEXT REFERENCES board_versions(id) NOT NULL,						
			name TEXT DEFAULT '',
			current_firware TEXT DEFAULT '',
			latest_firmware TEXT DEFAULT '' ,
			owner_id TEXT REFERENCES app_users(id) NOT NULL,
			serial_number TEXT DEFAULT '');`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS activation_card_types (
			id TEXT PRIMARY KEY,						
			color TEXT DEFAULT '',
			description TEXT DEFAULT '',
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS activation_cards (
			id TEXT PRIMARY KEY,	
			activation_card_type_id TEXT REFERENCES activation_card_types(id) NOT NULL,
			board_id TEXT REFERENCES boards(id),
			number TEXT DEFAULT '',
			code TEXT DEFAULT '',
			created_at  TIMESTAMP NOT NULL,
			updated_at  TIMESTAMP NOT NULL,			
			activated_by TEXT REFERENCES app_users(id) ,
			activated_at TIMESTAMP);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS waters (
			id TEXT PRIMARY KEY,	
			app_user_id TEXT REFERENCES app_users (id) NOT NULL,
			ext_id TEXT,
			name TEXT DEFAULT '',			
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS map_markers (
			id TEXT PRIMARY KEY,	
			app_user_id TEXT REFERENCES app_users (id) NOT NULL,
			ext_id TEXT DEFAULT '',
			board_id TEXT REFERENCES boards (id) NOT NULL,
			water_id TEXT REFERENCES waters (id) NOT NULL,
			type INTEGER DEFAULT 0,
			lat REAL NOT NULL,
			lng REAL NOT NULL,
			title TEXT NOT NULL,
			snippet TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS properties (
			id TEXT PRIMARY KEY,	
			type INTEGER NOT NULL,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS properties_trans (
			id TEXT PRIMARY KEY,	
			property_id TEXT REFERENCES properties (id) NOT NULL,
			lng_code TEXT NOT NULL,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS properties_list_values (
			id TEXT PRIMARY KEY,	
			property_id TEXT REFERENCES properties (id) NOT NULL,			
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS properties_list_values_trans (
			id TEXT PRIMARY KEY,	
			property_list_value_id TEXT REFERENCES properties_list_values (id) NOT NULL,			
			lng_code TEXT NOT NULL,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS board_version_properties (
			id TEXT PRIMARY KEY,
			board_version_id  TEXT REFERENCES board_versions (id) NOT NULL,			
			property_id TEXT REFERENCES properties (id) NOT NULL,			
			name TEXT NOT NULL,
			value_type INTEGER NOT NULL,
			value TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS board_settings (
			id TEXT PRIMARY KEY,
			board_id  TEXT REFERENCES boards (id) NOT NULL,			
			app_user_id TEXT REFERENCES app_users (id) NOT NULL,			
			property_id TEXT REFERENCES properties (id) NOT NULL,			
			name TEXT NOT NULL,
			value_type INTEGER NOT NULL,
			value TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS cloud_users (
	// 		id SERIAL PRIMARY KEY,
	// 		id_settings INT REFERENCES cloud_settings(id),
	// 		phone TEXT DEFAULT '',
	// 		pin TEXT DEFAULT '');`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS cloud_db_auth (
	// 		id SERIAL PRIMARY KEY,
	// 		id_cloud_db INT REFERENCES cloud_settings(id),
	// 		cloud_user TEXT,
	// 		cloud_password TEXT);`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
