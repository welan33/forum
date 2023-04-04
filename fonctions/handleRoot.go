package forum

import (
	"database/sql"
	"net/http"
	"text/template"
)

type Value struct {
	Id            int
	User_id       string
	Title         string
	Content       string
	Category      string
	Publication_d string
	Like          int
	Dislike       int
	Score         int
	Username      string
	Content_sum   string
	Image         string
	Logged        bool
}

type Full struct {
	Image       string
	Lien        string
	Class       string //= id
	Description string
	Email       string
	Rank        int
	Logged      bool
	Username    string
	Categories  []Categories
	Value       []Value
}

func MainHandler(res http.ResponseWriter, req *http.Request) {
	FiltreCat := ""
	FiltreLike := false
	FiltreOld := false
	if req.Method == "GET" {
		FiltreCat = ""
	} else if req.Method == "POST" {
		req.ParseForm()
		FiltreCat = req.Form.Get("cat")
		if req.Form.Get("like") == "like" {
			FiltreLike = true
		}
		if req.Form.Get("old") == "old" {
			FiltreOld = true
		}
	}
	db, _ := sql.Open("sqlite3", "forum.db")
	F := GetIndexData(res, req, db, FiltreCat, FiltreLike, FiltreOld)
	name_cat := ""
	id := 0
	datacat, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		panic(err)
	}
	defer datacat.Close()
	for datacat.Next() {
		err := datacat.Scan(&id, &name_cat)
		if err != nil {
			panic(err)
		}
		F.Categories = append(F.Categories, Categories{id, name_cat})
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(res, F)
}
