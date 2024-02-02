package db

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

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

	CREATE TABLE IF NOT EXISTS submission (
		id INTEGER PRIMARY KEY,
		date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		data JSON
	);
	`
	_, err := db.db.Exec(createTables)
	return err
}

func (db *DB) StoreState(state models.ClinicOutcomesFormState) error {
	b, err := json.Marshal(state)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO state (data) VALUES (?)", string(b))
	return err
}

func (db *DB) GetState(id string) (models.ClinicOutcomesFormState, error) {
	var state models.ClinicOutcomesFormState
	var b []byte
	err := db.db.QueryRow("SELECT data FROM state WHERE id = ?").Scan(&b)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(b, &state)
	return state, err
}

func (db *DB) StoreSubmission(submission models.ClinicOutcomesFormSubmit) error {
	b, err := json.Marshal(submission)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO submission (data) VALUES (?)", string(b))
	return err
}

type Submission struct {
	ID          string
	Data        models.ClinicOutcomesFormSubmit
	DateCreated time.Time
}

func (db *DB) GetSubmission(id string) (Submission, error) {
	var submission Submission
	var data []byte
	var dateCreated time.Time
	err := db.db.QueryRow("SELECT data,date_created FROM submission WHERE id = ?", id).Scan(&data, &dateCreated)
	if err != nil {
		return submission, err
	}
	var os models.ClinicOutcomesFormSubmit
	err = json.Unmarshal(data, &os)
	if err != nil {
		return submission, err
	}
	submission.ID = id
	submission.DateCreated = dateCreated
	submission.Data = os

	return submission, nil
}

func (db *DB) GetAllSubmissions() ([]Submission, error) {
	var submissions []Submission
	rows, err := db.db.Query("SELECT id,data,date_created FROM submission")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var data []byte
		var dateCreated time.Time
		err = rows.Scan(&id, &data, &dateCreated)
		if err != nil {
			return nil, err
		}
		var submission Submission
		var os models.ClinicOutcomesFormSubmit
		err = json.Unmarshal(data, &os)
		if err != nil {
			return nil, err
		}
		submission.ID = strconv.Itoa(id)
		submission.DateCreated = dateCreated
		submission.Data = os
		submissions = append(submissions, submission)
	}
	return submissions, rows.Err()
}
