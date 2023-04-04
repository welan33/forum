package forum

import (
	"database/sql"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func PostHandler(res http.ResponseWriter, req *http.Request) {
	//à modifier pour pouvoir ajouter ou supprimer des catégories

	db, _ := sql.Open("sqlite3", "forum.db")
	F := GetIndexData(res, req, db, "", false, false)
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
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("./static/poste.html"))
	tmpl.Execute(res, F)
}

func InsertPost(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.Redirect(res, req, "/", http.StatusFound)
	case "POST":
		req.ParseForm()
		title := req.Form.Get("title")
		content := req.Form.Get("content")
		if content == "" {
			http.Redirect(res, req, "/poste", http.StatusFound)
			return
		}
		category_id := req.Form.Get("category")
		image := req.Form.Get("image")
		println(image)
		like := 0
		dislike := 0
		score := like - dislike
		date := time.Now().Format("2006-01-02 15:04:05")
		//récupération de l'user id via le cookie
		cookie, _ := req.Cookie("Jbzz")
		if cookie.Value == "" {
			http.Redirect(res, req, "/co", http.StatusFound)
		} else {
			db, err := sql.Open("sqlite3", "forum.db")
			if err != nil {
				panic(err)
			}
			_, err = db.Exec(`INSERT INTO postes (user_id, title, like, dislike, score, publication_d, category_id, content, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`, cookie.Value, title, like, dislike, score, date, category_id, content, image)
			if err != nil {
				panic(err)
			}
			//rediriger vers la page fraichement créée du poste
			defer db.Close()
			http.Redirect(res, req, "/", http.StatusFound)
			return
		}
		http.Redirect(res, req, "/co", http.StatusFound)
	}
}

func LikeHandler(res http.ResponseWriter, req *http.Request) {
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	db, _ := sql.Open("sqlite3", "forum.db")
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	if AddPostLikeToUser(idpath, cookie.Value, db) {
		_, err := db.Exec(`UPDATE postes SET like = like + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(`UPDATE postes SET score = score + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(`UPDATE postes SET like = like - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(`UPDATE postes SET score = score - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}

func DislikeHandler(res http.ResponseWriter, req *http.Request) {
	idpath := req.URL.Path
	idpath = strings.Split(idpath, "/")[len(strings.Split(idpath, "/"))-1]
	db, _ := sql.Open("sqlite3", "forum.db")
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		http.Redirect(res, req, "/co", http.StatusFound)
		defer db.Close()
		return
	}
	if AddPostDislikeToUser(idpath, cookie.Value, db) {
		_, err := db.Exec(`UPDATE postes SET dislike = dislike + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(`UPDATE postes SET score = score - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := db.Exec(`UPDATE postes SET dislike = dislike - 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(`UPDATE postes SET score = score + 1 WHERE id=?;`, idpath)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
	http.Redirect(res, req, "/", http.StatusFound)
}
