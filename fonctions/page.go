package forum

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"
)

func PageDisplay(res http.ResponseWriter, req *http.Request) {

	type Commentaire struct {
		Id            int
		Like          int
		Dislike       int
		Poste_id      int
		Comment       string
		User_id       string
		Publication_d string
		ComUsername   string
	}

	type User struct {
		Lien       string
		Logged     bool
		UsUsername string
	}

	type Post struct {
		Id            int
		Like          int
		Dislike       int
		Score         int
		User_id       string
		PostUsername  string
		Title         string
		Content       string
		Category      string
		Publication_d string
		Image         string
		Commentaires  []Commentaire
		User          User
	}

	var P = Post{}
	post_id := ""

	//si clic sur le titre ou accès via url redirect
	path := req.URL.Path
	path = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	post_id = path
	//si clic via bouton accès au poste
	req.ParseForm()
	if req.Form.Get("id") != "" {
		post_id = req.Form.Get("id")
	}
	db, _ := sql.Open("sqlite3", "forum.db")
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || cookie.Value == "" {
		P.User.Logged = false
		P.User.Lien = "/co"
	} else {
		P.User.Logged = true
		rows, _ := db.Query("SELECT username FROM user WHERE id = ?", cookie.Value)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&P.User.UsUsername)
			if err != nil {
				panic(err)
			}
		}
		P.User.Lien = "/profil" + P.User.UsUsername
	}
	cat_id := 0
	result, _ := db.Query(`SELECT * FROM postes WHERE id = ?;`, post_id)
	defer result.Close()
	for result.Next() {
		err := result.Scan(&P.Id, &P.User_id, &P.Title, &P.Content, &P.Like, &P.Dislike, &P.Score, &P.Publication_d, &cat_id, &P.Image)
		if err != nil {
			panic(err)
		}
	}
	datacat, _ := db.Query("SELECT name FROM categories WHERE id = ?", cat_id)
	defer datacat.Close()
	for datacat.Next() {
		err := datacat.Scan(&P.Category)
		if err != nil {
			panic(err)
		}
	}
	defer result.Close()
	result, _ = db.Query(`SELECT username FROM user WHERE id = ?;`, P.User_id)
	defer result.Close()
	for result.Next() {
		err := result.Scan(&P.PostUsername)
		if err != nil {
			panic(err)
		}
	}
	result, _ = db.Query(`SELECT * FROM comments WHERE poste_id = ?;`, post_id)
	var comment, user_id, publication_d, username string
	var id, like, dislike, poste_id int
	defer result.Close()
	for result.Next() {
		err := result.Scan(&id, &comment, &user_id, &poste_id, &publication_d, &like, &dislike)
		if err != nil {
			panic(err)
		}
		com_username, _ := db.Query(`SELECT username FROM user WHERE id = ?;`, user_id)
		defer com_username.Close()
		for com_username.Next() {
			err := com_username.Scan(&username)
			if err != nil {
				panic(err)
			}
		}

		P.Commentaires = append(P.Commentaires, Commentaire{id, like, dislike, poste_id, comment, user_id, publication_d, username})
	}

	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./static/page.html"))
	tmpl.Execute(res, P)
}
