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
		date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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

type Save struct {
	ID          string
	Data        models.ClinicOutcomesFormPayload
	DateCreated time.Time
}

func (db *DB) StoreState(state models.ClinicOutcomesFormPayload) error {
	b, err := json.Marshal(state)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO state (data) VALUES (?)", string(b))
	return err
}

func (db *DB) UpdateState(id string, newState models.ClinicOutcomesFormPayload) error {
	b, err := json.Marshal(newState)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("UPDATE state SET data = ? WHERE id = ?", string(b), id)
	return err
}

func (db *DB) GetState(id string) (Save, error) {
	var save Save
	var b []byte
	var dateCreated time.Time
	err := db.db.QueryRow("SELECT data, date_created FROM state WHERE id = ?", id).Scan(&b, &dateCreated)
	if err != nil {
		return save, err
	}
	var state models.ClinicOutcomesFormPayload
	err = json.Unmarshal(b, &state)
	if err != nil {
		return save, err
	}
	save.ID = id
	save.DateCreated = dateCreated
	save.Data = state
	return save, nil
}

func (db *DB) GetAllStates() ([]Save, error) {
	var saves []Save
	rows, err := db.db.Query("SELECT id,data,date_created FROM state")
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
		var save Save
		var state models.ClinicOutcomesFormPayload
		err = json.Unmarshal(data, &state)
		if err != nil {
			return nil, err
		}
		save.ID = strconv.Itoa(id)
		save.DateCreated = dateCreated
		save.Data = state
		saves = append(saves, save)
	}
	return saves, rows.Err()
}

type Submission struct {
	ID          string
	Data        models.ClinicOutcomesFormSubmit
	DateCreated time.Time
}

func (db *DB) StoreSubmission(submission models.ClinicOutcomesFormSubmit) error {
	b, err := json.Marshal(submission)
	if err != nil {
		return err
	}
	_, err = db.db.Exec("INSERT INTO submission (data) VALUES (?)", string(b))
	return err
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
	// Ticket 51

	// Get all submissions data from the database
	rows, err := db.db.Query("select id,data,date_created FROM submission")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Submissions slice to hold data from the returned rows
	var submissions []Submission

	// Iterate over the rows and append the data to the slice of Submissions
	for rows.Next() {
		var id string
		var formData []byte
		var dateCreated time.Time

		//scan the row
		err := rows.Scan(&id, &formData, &dateCreated)
		if err != nil {
			return nil, err
		}

		var submission Submission
		var os models.ClinicOutcomesFormSubmit
		err = json.Unmarshal(formData, &os)
		if err != nil {
			return nil, err
		}

		submission.ID = id
		submission.DateCreated = dateCreated
		submission.Data = os
		submissions = append(submissions, submission)
	}
	return submissions, rows.Err()

	// Refer to the GetAllStates function above for an example of how to do this
}
