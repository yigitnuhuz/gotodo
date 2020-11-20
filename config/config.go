package config

import "database/sql"

func GetDb() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./db/gotodo.db")

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Todos (Id INTEGER PRIMARY KEY, Detail TEXT, Completed BIT);")
	statement.Exec()
	return
}

const JwtTokenSecret = "GoToDo22"
