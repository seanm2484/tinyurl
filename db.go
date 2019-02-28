package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//URLRecord is the database structure for mapping longURL -> shortURL
type URLRecord struct {
	id       int
	longURL  string
	shortURL string
}

// initalize the sqlite db and return the object
func initDb(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func createTable(db *sql.DB) {
	// create table if not exists
	table := `
	CREATE TABLE IF NOT EXISTS records(
		Id INTEGER NOT NULL PRIMARY KEY,
		longURL TEXT,
		shortURL TEXT,
		InsertedDatetime DATETIME
	);
	`
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}

func writeRecord(db *sql.DB, url URLRecord) {
	addItemQuery := `
	INSERT OR REPLACE INTO records(
		Id,
		longURL,
		shortURL,
		InsertedDatetime
	) values(?, ?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(addItemQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(url.id, url.longURL, url.shortURL)
	if err2 != nil {
		panic(err2)
	}
}

func readItem(db *sql.DB, shortURL string) URLRecord {
	/*readAllQuery := `
	SELECT Id, longURL, shortURL FROM records
	ORDER BY datetime(InsertedDatetime) DESC
	`*/
	readAllQuery := `
        SELECT Id, longURL FROM records
        WHERE shortURL = (?)
    `
	rows, err := db.Query(readAllQuery, shortURL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result URLRecord
	for rows.Next() {
		item := URLRecord{}
		err2 := rows.Scan(&item.id, &item.longURL)
		if err2 != nil {
			panic(err2)
		}
		//result = append(result, item)
		return item
	}
	//return item
	return result
}
