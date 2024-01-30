package db

import (
	"database/sql"
	"encoding/json"

	"github.com/atos-digital/DHCW-clinic-outcomes/internal/server/models"
)

type DB struct {
	db *sql.DB
}

func NewSqlite(path string) (*DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Migrate() error {
	const createTables = `
	CREATE TABLE IF NOT EXISTS outcomes (
		id INTEGER PRIMARY KEY,
		state JSON
	);
	`
	_, err := db.db.Exec(createTables)
	return err
}

func (db *DB) StoreState(state models.OutcomesState) error {
	b, err := json.Marshal(state)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO outcomes (state) VALUES (?)", string(b))
	return err
}

func (db *DB) GetState(id string) (models.OutcomesState, error) {
	var state models.OutcomesState
	var b []byte
	err := db.db.QueryRow("SELECT state FROM outcomes WHERE id = ?").Scan(&b)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(b, &state)
	return state, err
}
