package forum

import (
	"database/sql"
	"net/http"
)

func CoHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/login.html")
}

func LogHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/co", http.StatusFound)
	case "POST":
		req.ParseForm()
		email := req.Form.Get("email")
		password := req.Form.Get("password")
		db, _ := sql.Open("sqlite3", "forum.db")
		result, _ := db.Query(`SELECT password, id FROM user WHERE email = ?;`, email)
		defer db.Close()
		var hash, uuid string
		defer result.Close()
		for result.Next() {
			err := result.Scan(&hash, &uuid)
			if err != nil {
				panic(err)
			}
		}
		if CheckPasswordHash(password, hash) {
			http.SetCookie(res, &http.Cookie{Name: "Jbzz", Value: uuid})
			http.Redirect(res, req, "/", http.StatusFound)
		} else {
			http.Redirect(res, req, "/co", http.StatusFound)
		}
	}
}

func OutHandler(res http.ResponseWriter, req *http.Request) {
	//reset les cookies
	http.SetCookie(res, &http.Cookie{Name: "Jbzz", MaxAge: -1})
	http.Redirect(res, req, "/", http.StatusFound)
}

func GoogleLogHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/co", http.StatusFound)
	case "POST":
		req.ParseForm()
		email := req.Form.Get("email")
		id := req.Form.Get("id")
		db, _ := sql.Open("sqlite3", "forum.db")
		result, _ := db.Query(`SELECT id FROM user WHERE email = ?;`, email)
		defer db.Close()
		var uuid string //version hash de l'id
		defer result.Close()
		for result.Next() {
			err := result.Scan(&uuid)
			if err != nil {
				panic(err)
			}
		}
		if CheckPasswordHash(id, uuid) {
			http.SetCookie(res, &http.Cookie{Name: "Jbzz", Value: uuid})
			http.Redirect(res, req, "/", http.StatusFound)
		} else {
			http.Redirect(res, req, "/co", http.StatusFound)
		}
	}
}
