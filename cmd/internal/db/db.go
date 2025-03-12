package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	DB *sql.DB
}

func New() (*DataBase, error) {
	var DB *sql.DB
	var err error

	DB, err = sql.Open("sqlite3", "./app.db")

	return &DataBase{DB}, err

}

func (db *DataBase) Init() error {

	companyStmt := `
	CREATE TABLE IF NOT EXISTS company(
	id TEXT NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	clientStmt := `
	CREATE TABLE IF NOT EXISTS client(
	id TEXT NOT NULL PRIMARY KEY, 
	name TEXT NOT NULL,
	dob TEXT NOT NULL,
	phone TEXT NOT NULL,
	email TEXT NOT NULL,
	address TEXT NOT NULL,
	company_id TEXT NOT NULL, 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
	);
	`
	serviceStmt := `
	CREATE TABLE IF NOT EXISTS service(
	id TEXT NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	due_days INTEGER NOT NULL,
	priority INTEGER NOT NULL,
	max_reschedules INTEGER NOT NULL,
	company_id TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
	);
	`

	_, err := db.DB.Exec(companyStmt)
	if err != nil {
		log.Fatalf("Error creating the company table %q", err)
		return err
	}

	_, err1 := db.DB.Exec(clientStmt)
	if err1 != nil {
		log.Fatalf("Error creating the clients table %q", err1)
		return err1
	}

	_, err2 := db.DB.Exec(serviceStmt)
	if err2 != nil {
		log.Fatalf("Error creating the service table %q", err2)
		return err2
	}
	return nil

}
