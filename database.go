package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// InitializeDB inititalizes database
func InitializeDB() {
	database, _ := sql.Open("sqlite3", "./emans.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, eid INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO employees (firstname, lastname, eid) VALUES (?, ?, ?)")
	statement.Exec("James", "Bond", "007")
	rows, _ := database.Query("SELECT id, firstname, lastname, eid FROM employees")
	var id int
	var firstname string
	var lastname string
	var eid int
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &eid)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + strconv.Itoa(eid))
	}
}
