package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		panic("Error" + err.Error())
	}
	return database
}

func Migrate() {
	db := GetDB()

	create_table_query := `CREATE TABLE notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    content TEXT
);`

	statement, _ := db.Prepare(create_table_query)
	statement.Exec()

}
