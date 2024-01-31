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
	CREATE TABLE IF NOT EXISTS state (
		id INTEGER PRIMARY KEY,
		data JSON
	);

	CREATE TABLE IF NOT EXISTS submissions (
		id INTEGER PRIMARY KEY,
		data JSON
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
	_, err = db.db.Exec("INSERT INTO state (data) VALUES (?)", string(b))
	return err
}

func (db *DB) GetState(id string) (models.OutcomesState, error) {
	var state models.OutcomesState
	var b []byte
	err := db.db.QueryRow("SELECT data FROM state WHERE id = ?").Scan(&b)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(b, &state)
	return state, err
}

func (db *DB) StoreSubmission(submission models.OutcomesSubmit) error {
	b, err := json.Marshal(submission)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO submissions (data) VALUES (?)", string(b))
	return err
}

func (db *DB) GetSubmission(id string) (models.OutcomesSubmit, error) {
	var submission models.OutcomesSubmit
	var b []byte
	err := db.db.QueryRow("SELECT data FROM submissions WHERE id = ?").Scan(&b)
	if err != nil {
		return submission, err
	}
	err = json.Unmarshal(b, &submission)
	return submission, err
}
