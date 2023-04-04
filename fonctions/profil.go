package forum

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"
)

func ProfilHandler(res http.ResponseWriter, req *http.Request) {

	//si clic sur le pseudo ou accès via url direct
	path := req.URL.Path
	path = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]

	//si clic via bouton accès au profil
	req.ParseForm()
	if req.Form.Get("username") != "" {
		path = req.Form.Get("username")
	}

	db, _ := sql.Open("sqlite3", "forum.db")
	var U = GetIndexData(res, req, db, "", false, false)
	result, _ := db.Query("SELECT email, username, picture, description, rank, id FROM user WHERE username = ?", path)
	var id string
	if result != nil {
		defer result.Close()
		for result.Next() {
			err := result.Scan(&U.Email, &U.Username, &U.Image, &U.Description, &U.Rank, &id)
			if err != nil {
				panic(err)
			}
		}
	}
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || cookie.Value == "" {
		U.Logged = false
	} else {
		rows, _ := db.Query("SELECT username FROM user WHERE id = ?", cookie.Value)
		defer rows.Close()
		var checkname string
		for rows.Next() {
			err := rows.Scan(&checkname)
			if err != nil {
				panic(err)
			}
		}
		if checkname == path {
			U.Logged = true
		} else {
			U.Logged = false
		}
	}
	defer db.Close()
	template := template.Must(template.ParseFiles("static/profil.html"))
	template.Execute(res, U)
}

func ModifProfile(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		email := req.Form.Get("email")
		description := req.Form.Get("description")
		image := req.Form.Get("image")
		db, _ := sql.Open("sqlite3", "forum.db")
		cookie, err := req.Cookie("Jbzz")
		if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
			http.Redirect(res, req, "/co", http.StatusFound)
			defer db.Close()
			return
		}
		if username == "" || email == "" {
			rows, _ := db.Query("SELECT username FROM user WHERE username = ?", username)
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&username)
			}
			http.Redirect(res, req, "/profil"+username, http.StatusFound)
			return
		}
		_, err = db.Exec(`UPDATE user SET username = ?, email = ?, description = ?, picture = ? WHERE id = ?;`, username, email, description, image, cookie.Value)
		if err != nil {
			panic(err)
		}
		http.Redirect(res, req, "/profil/"+username, http.StatusFound)
		defer db.Close()
		return
	}
}
