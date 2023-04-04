package forum

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	//ne pas fermer la db ici
	return db
}

func CreateTable(db *sql.DB) {
	tables := `
    CREATE TABLE IF NOT EXISTS user (
    	id TEXT PRIMARY KEY,
		email TEXT,
		username TEXT,
		password TEXT,
		picture TEXT,
		description TEXT,
		like TEXT,
		rank INT
		);

	CREATE TABLE IF NOT EXISTS postes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		title TEXT,
		content TEXT,
		like INT,
		dislike INT,
		score INT,
		publication_d TEXT,
		category_id  INT,
		image TEXT,
		FOREIGN KEY(user_id) REFERENCES user(id)
		FOREIGN KEY(category_id) REFERENCES categories(id)
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		comment TEXT,
		user_id INT,
		poste_id INT,
		publication_d TEXT,
		like INT,
		dislike INT,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(poste_id) REFERENCES postes(id)
	);

	CREATE TABLE IF NOT EXISTS requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		modoname TEXT,
		user_id INT,
		publication_d TEXT,
		FOREIGN KEY(user_id) REFERENCES user(id)
	);

	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	`
	_, err := db.Exec(tables)
	if err != nil {
		panic(err)
	}
}

func AlterDB() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`ALTER TABLE user DROP COLUMN firstname;`)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func DeletepostDB(res http.ResponseWriter, req *http.Request) {
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`DELETE FROM postes WHERE id=?;`, idpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}

func DeletecommentDB(res http.ResponseWriter, req *http.Request) {
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`DELETE FROM comments WHERE id=?;`, idpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}

func DeleteuserDB(res http.ResponseWriter, req *http.Request) {
	userpath := req.URL.Path
	userpath = strings.Split(userpath, "/")[len(strings.Split(userpath, "/"))-1]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`DELETE FROM user WHERE id=?;`, userpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}

func DeletecatDB(res http.ResponseWriter, req *http.Request) {
	userpath := req.URL.Path
	userpath = strings.Split(userpath, "/")[len(strings.Split(userpath, "/"))-1]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`DELETE FROM categories WHERE id=?;`, userpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}

func DeleterequestDB(res http.ResponseWriter, req *http.Request) {
	userpath := req.URL.Path
	userpath = strings.Split(userpath, "/")[len(strings.Split(userpath, "/"))-1]
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic("failed to connect database")
	}
	_, err = db.Exec(`DELETE FROM requests WHERE id=?;`, userpath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}
