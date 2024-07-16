package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the Database.")
	}

	// 10 conns opened simultaneously at most
	// then we have a pool of 10 conns that can be used
	// should not also be too big
	// if we wanted to make more than 10 conns, these other conns will wait until there are available ones
	DB.SetMaxOpenConns(10)

	// how many conns we want to keep open if no one using these conns at the moment
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println("fy haga hnaa 8ltt")
		// panic("Could not create events table.")
	}
}
