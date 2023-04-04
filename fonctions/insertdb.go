package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InsertUserDB(username string, email string, password string, uuid string) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}

	_, err = db.Exec(`INSERT INTO user (username, email, password, id, picture, description, like, rank) VALUES (?, ?, ?, ?, "", "", "", 0)`, username, email, password, uuid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func InsertGUserDB(username string, email string, picture string, uuid string) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}

	_, err = db.Exec(`INSERT INTO user (username, email, password, id, picture, description, like, rank) VALUES (?, ?, "", ?, ?, "", "", 0)`, username, email, uuid, picture)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
