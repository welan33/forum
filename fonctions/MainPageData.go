package forum

import (
	"database/sql"
	"net/http"
)

func GetIndexData(res http.ResponseWriter, req *http.Request, db *sql.DB, filtreCat string, SortLike bool, SortOld bool) Full {
	var F = Full{}

	result, _ := db.Query("SELECT * FROM postes")
	var title, user_id, content, category, publication_d, username, content_sum, image string
	var id, like, dislike, score, category_id int
	cookie, err := req.Cookie("Jbzz")
	if err != nil || cookie == nil || !DoUserExist(cookie.Value, db) {
		if cookie != nil {
			http.SetCookie(res, &http.Cookie{Name: "Jbzz", MaxAge: -1})
		}
		F.Logged = false
		F.Username = ""
		F.Image = "Se connecter"
		F.Lien = "/co"
	} else {
		F.Logged = true
		rows, _ := db.Query("SELECT username FROM user WHERE id = ?", cookie.Value)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&F.Username)
			if err != nil {
				panic(err)
			}
		}
		F.Image = "Profil"
		F.Lien = "/profil/" + F.Username
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&id, &user_id, &title, &content, &like, &dislike, &score, &publication_d, &category_id, &image)
		if err != nil {
			panic(err)
		}
		datacat, err := db.Query("SELECT name FROM categories WHERE id = ?", category_id)
		if err != nil {
			panic(err)
		}
		defer datacat.Close()
		for datacat.Next() {
			err := datacat.Scan(&category)
			if err != nil {
				panic(err)
			}
		}
		dataname, err := db.Query("SELECT username FROM user WHERE id = ?", user_id)
		if err != nil {
			panic(err)
		}
		defer dataname.Close()
		for dataname.Next() {
			err := dataname.Scan(&username)
			if err != nil {
				panic(err)
			}
		}
		if len(content) > 400 {
			content_sum = content[:400]
			for i := range content_sum {
				if content_sum[len(content_sum)-1-i] == ' ' {
					content_sum = content_sum[:len(content_sum)-1-i] + "..."
					break
				}
			}
			if content_sum[len(content_sum)-3:] != "..." {
				content_sum += "..."
			}
		} else {
			content_sum = content
		}
		log := false
		if F.Logged && user_id == cookie.Value {
			log = true
		}
		if filtreCat == "" || category == filtreCat {
			F.Value = append(F.Value, Value{id, user_id, title, content, category, publication_d, like, dislike, score, username, content_sum, image, log})
		}
	}
	if SortLike {
		F.Value = SortByLike(F.Value)
	} else if SortOld {
		F.Value = SortByOld(F.Value)
	}
	return F
}
