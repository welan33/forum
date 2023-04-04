package forum

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func CommentHandler(res http.ResponseWriter, req *http.Request) {
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
	tmpl := template.Must(template.ParseFiles("./static/comment.html"))
	tmpl.Execute(res, F)
}

func InsertComment(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		content := req.Form.Get("content")
		post_id := req.Form.Get("id")
		//postuser := req.Form.Get("PostUsername")
		if content == "" {
			println("empty content")
			http.Redirect(res, req, "/comment", http.StatusFound)
		}
		like := 0
		dislike := 0
		date := time.Now().Format("2006-01-02 15:04:05")

		//récupération de l'user id via le cookie
		db, _ := sql.Open("sqlite3", "forum.db")
		cookie, err := req.Cookie("Jbzz")
		if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
			panic(err)
		} else {
			_, err = db.Exec(`INSERT INTO comments (user_id, poste_id, like, dislike, publication_d, comment) VALUES (?, ?, ?, ?, ?, ?);`, cookie.Value, post_id, like, dislike, date, content)
			if err != nil {
				panic(err)
			}
		}
		defer db.Close()
		http.Redirect(res, req, "/page/"+post_id, http.StatusFound)
	}
}

func LikeCHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	postid := req.Form.Get("id")
	db, _ := sql.Open("sqlite3", "forum.db")
	user_id, _ := db.Query(`SELECT user_id FROM postes WHERE id = ?;`, postid)
	var id_string string
	defer user_id.Close()
	for user_id.Next() {
		user_id.Scan(&id_string)
	}
	username, _ := db.Query(`SELECT username FROM user WHERE id = ?;`, id_string)
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	if AddComLikeToUser(idpath, cookie.Value, db) {
		_, err := db.Exec(`UPDATE comments SET like = like + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(`UPDATE comments SET like = like - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	}
	var username_string string
	defer username.Close()
	for username.Next() {
		username.Scan(&username_string)
	}
	defer db.Close()
	http.Redirect(res, req, "/page/"+postid, http.StatusFound)
}

func DislikeCHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	postid := req.Form.Get("id")
	//postuser := req.Form.Get("PostUsername")
	db, _ := sql.Open("sqlite3", "forum.db")
	user_id, _ := db.Query(`SELECT user_id FROM postes WHERE id = ?;`, postid)
	var id_string string
	defer user_id.Close()
	for user_id.Next() {
		user_id.Scan(&id_string)
	}
	username, _ := db.Query(`SELECT username FROM user WHERE id = ?;`, id_string)
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	if AddComDislikeToUser(idpath, cookie.Value, db) {
		_, err := db.Exec(`UPDATE comments SET dislike = dislike + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(`UPDATE comments SET dislike = dislike - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	}
	var username_string string
	defer username.Close()
	for username.Next() {
		username.Scan(&username_string)
	}
	defer db.Close()
	http.Redirect(res, req, "/page/"+postid, http.StatusFound)
}
