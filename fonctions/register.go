package forum

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	_ "golang.org/x/crypto/bcrypt"
)

func RegHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("username")
		email := req.Form.Get("email")
		password := req.Form.Get("password")
		password = hpwd(password)
		uuid := uuid.New().String()
		if password == "" {
			http.Redirect(res, req, "/co", http.StatusFound)
		}
		db, _ := sql.Open("sqlite3", "forum.db")
		if !DoEmailExist(email, db) && !DoUsernameExist(username, db) {
			InsertUserDB(username, email, password, uuid)
		} else {
			http.Redirect(res, req, "/co", http.StatusFound)
			return
		}
		defer db.Close()
		http.SetCookie(res, &http.Cookie{Name: "Jbzz", Value: uuid})
		username, email, password, uuid = "", "", "", ""
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func GoogleHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		username := req.Form.Get("name")
		image := req.Form.Get("image")
		email := req.Form.Get("email")
		uuid := hpwd(req.Form.Get("id"))
		db, _ := sql.Open("sqlite3", "forum.db")
		if !DoEmailExist(email, db) && !DoUsernameExist(username, db) {
			InsertGUserDB(username, email, image, uuid)
		} else {
			http.Redirect(res, req, "/co", http.StatusFound)
			return
		}
		defer db.Close()
		http.SetCookie(res, &http.Cookie{Name: "Jbzz", Value: uuid})
		username, email, uuid = "", "", ""
		http.Redirect(res, req, "/", http.StatusFound)
	}
}