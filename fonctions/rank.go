package forum

import (
	"database/sql"
	"net/http"
	"text/template"
)

func UserToAdmin(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			panic(err)
		}
		rows, _ := db.Query(`SELECT id FROM user WHERE username = ?;`, username)
		defer rows.Close()
		var user_id string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&user_id)
		}
		_, err = db.Exec(`UPDATE user SET rank = 3 WHERE id = ?;`, user_id)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		println("Rank modifié")
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func UserToModerator(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			panic(err)
		}
		rows, _ := db.Query(`SELECT id FROM user WHERE username = ?;`, username)
		defer rows.Close()
		var user_id string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&user_id)
		}
		_, err = db.Exec(`UPDATE user SET rank = 2 WHERE id = ?;`, user_id)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func VIPHandler(res http.ResponseWriter, req *http.Request) {
	db, _ := sql.Open("sqlite3", "forum.db")
	F := GetIndexData(res, req, db, "", false, false)
	tmpl := template.Must(template.ParseFiles("./static/vip.html"))
	tmpl.Execute(res, F)
}

func UserToVIP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			panic(err)
		}
		rows, _ := db.Query(`SELECT id FROM user WHERE username = ?;`, username)
		defer rows.Close()
		var user_id string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&user_id)
		}
		_, err = db.Exec(`UPDATE user SET rank = 1 WHERE id = ?;`, user_id)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func ResetRank(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			panic(err)
		}
		rows, _ := db.Query(`SELECT id FROM user WHERE username = ?;`, username)
		defer rows.Close()
		var user_id string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&user_id)
		}
		_, err = db.Exec(`UPDATE user SET rank = 0 WHERE id = ?;`, user_id)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(res, req, "/", http.StatusFound)
	}
}
