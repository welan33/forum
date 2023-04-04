package forum

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Users struct {
	Id       string
	Rank     int
	Username string
}

type Request struct {
	ModoName string
	Id       int
	Content  string
	Date     string
}

type Categories struct {
	Id   int
	Name string
}
type Admin struct {
	User_id    string
	Rank       int
	Categories []Categories
	UserList   []Users
	Request    []Request
}

func AdminHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		A := Admin{}
		id_user := ""
		rank := 0
		username := ""
		id_cat := 0
		name_cat := ""
		id_req := 0
		modoname := ""
		content := ""
		date := ""
		path := req.URL.Path
		user := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
		db, _ := sql.Open("sqlite3", "forum.db")
		result, _ := db.Query("SELECT id, rank FROM user WHERE username = ?", user)
		if result != nil {
			defer result.Close()
			for result.Next() {
				err := result.Scan(&A.User_id, &A.Rank)
				if err != nil {
					panic(err)
				}
			}
		} else {
			http.Redirect(res, req, "/", http.StatusFound)
		}
		if A.Rank == 3 {
			datauser, err := db.Query("SELECT id, rank, username FROM user")
			if err != nil {
				panic(err)
			}
			defer datauser.Close()
			for datauser.Next() {
				err := datauser.Scan(&id_user, &rank, &username)
				if err != nil {
					panic(err)
				}
				A.UserList = append(A.UserList, Users{id_user, rank, username})
			}
			datacat, err := db.Query("SELECT id, name FROM categories")
			if err != nil {
				panic(err)
			}
			defer datacat.Close()
			for datacat.Next() {
				err := datacat.Scan(&id_cat, &name_cat)
				if err != nil {
					panic(err)
				}
				A.Categories = append(A.Categories, Categories{id_cat, name_cat})
			}
			datareq, err := db.Query("SELECT id, modoname, content, publication_d FROM requests")
			if err != nil {
				panic(err)
			}
			defer datareq.Close()
			for datareq.Next() {
				err := datareq.Scan(&id_req, &modoname, &content, &date)
				if err != nil {
					panic(err)
				}
				A.Request = append(A.Request, Request{modoname, id_req, content, date})
			}
		} else {
			http.Redirect(res, req, "/", http.StatusFound)
		}
		defer db.Close()
		tmpl := template.Must(template.ParseFiles("./static/admin.html"))
		tmpl.Execute(res, A)
	}

}

func UpdateCategories(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		name := req.Form.Get("cat")
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			panic("failed to connect database")
		}

		_, err = db.Exec(`INSERT INTO categories (name) VALUES (?)`, name)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		name = ""
		http.Redirect(res, req, "/", http.StatusFound)
	}
}

func RequestHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	db, _ := sql.Open("sqlite3", "forum.db")
	F := GetIndexData(res, req, db, "", false, false)
	F.Class = req.Form.Get("id")
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./static/request.html"))
	tmpl.Execute(res, F)
}

func SaveRequest(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		content := req.Form.Get("content")
		if content == "" {
			println("empty content")
			http.Redirect(res, req, "/request", http.StatusFound)
		}
		modo := req.Form.Get("PostUsername")
		date := time.Now().Format("2006-01-02 15:04:05")
		//récupération de l'user id via le cookie
		db, _ := sql.Open("sqlite3", "forum.db")
		cookie, err := req.Cookie("Jbzz")
		if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
			panic(err)
		} else {
			_, err = db.Exec(`INSERT INTO requests (content, modoname, user_id, publication_d) VALUES (?, ?, ?, ?);`, content, modo, cookie.Value, date)
			if err != nil {
				panic(err)
			}
		}
		defer db.Close()
		http.Redirect(res, req, "/", http.StatusFound)
	}
}
