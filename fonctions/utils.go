package forum

import (
	"database/sql"
	"sort"
	"strings"
)

func DoUserExist(id string, db *sql.DB) bool {
	rows, err := db.Query(`SELECT username FROM user WHERE id = ?;`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

func DoEmailExist(email string, db *sql.DB) bool {
	rows, err := db.Query(`SELECT email FROM user WHERE email = ?;`, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

func DoUsernameExist(username string, db *sql.DB) bool {
	rows, err := db.Query(`SELECT username FROM user WHERE username = ?;`, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		return true
	}
	return false
}

func AddPostLikeToUser(id_post string, id_user string, db *sql.DB) bool {
	rows, _ := db.Query(`SELECT like FROM user WHERE id = ?;`, id_user)
	defer rows.Close()
	var like string
	for rows.Next() {
		rows.Scan(&like)
	}
	likes := strings.Split(like, " ")
	var likesid [][]string
	for _, like := range likes {
		likesid = append(likesid, strings.Split(like, "="))
	}
	for _, value := range likesid {
		if value[0] == "post:"+id_post {
			if value[1] == "1" {
				value[1] = "0"
				like = ""
				for _, nval := range likesid {
					if nval[0] != "" {
						like += nval[0] + "=" + nval[1] + " " //coucou
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
				return false
			} else if value[1] == "-1" {
				_, _ = db.Exec(`UPDATE postes SET score = score + 1 WHERE id=?;`, id_post)
			}
			value[1] = "1"
			like = ""
			for _, nval := range likesid {
				if nval[0] != "" {
					like += nval[0] + "=" + nval[1] + " "
				}
			}
			_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
			return true
		}
	}
	like += " post:" + id_post + "=1"
	_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
	return true
}

func AddComLikeToUser(id_com string, id_user string, db *sql.DB) bool {
	rows, _ := db.Query(`SELECT like FROM user WHERE id = ?;`, id_user)
	defer rows.Close()
	var like string
	for rows.Next() {
		rows.Scan(&like)
	}
	likes := strings.Split(like, " ")
	var likesid [][]string
	for _, like := range likes {
		likesid = append(likesid, strings.Split(like, "="))
	}
	for _, value := range likesid {
		if value[0] == "comm:"+id_com {
			if value[1] == "1" {
				value[1] = "0"
				like = ""
				for _, nval := range likesid {
					if nval[0] != "" {
						like += nval[0] + "=" + nval[1] + " "
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
				return false
			} else {
				value[1] = "1"
				like = ""
				for _, nval := range likesid {
					if nval[0] != "" {
						like += nval[0] + "=" + nval[1] + " "
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
				return true
			}
		}
	}
	like += " comm:" + id_com + "=1"
	_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, like, id_user)
	return true
}

func AddComDislikeToUser(id_com string, id_user string, db *sql.DB) bool {
	rows, _ := db.Query(`SELECT like FROM user WHERE id = ?;`, id_user)
	defer rows.Close()
	var dislike string
	for rows.Next() {
		rows.Scan(&dislike)
	}
	dislikes := strings.Split(dislike, " ")
	var dislikesid [][]string
	for _, dislike := range dislikes {
		dislikesid = append(dislikesid, strings.Split(dislike, "="))
	}
	for _, value := range dislikesid {
		if value[0] == "comm:"+id_com {
			if value[1] == "-1" {
				value[1] = "0"
				dislike = ""
				for _, nval := range dislikesid {
					if nval[0] != "" {
						dislike += nval[0] + "=" + nval[1] + " "
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, dislike, id_user)
				return false
			} else {
				value[1] = "-1"
				dislike = ""
				for _, nval := range dislikesid {
					if nval[0] != "" {
						dislike += nval[0] + "=" + nval[1] + " "
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, dislike, id_user)
				return true
			}
		}
	}
	dislike += " comm:" + id_com + "=-1"
	_, _ = db.Exec(`UPDATE user SET dislike = ? WHERE id = ?;`, dislike, id_user)
	return true
}

func AddPostDislikeToUser(id_post string, id_user string, db *sql.DB) bool {
	rows, _ := db.Query(`SELECT like FROM user WHERE id = ?;`, id_user)
	defer rows.Close()
	var dislike string
	for rows.Next() {
		rows.Scan(&dislike)
	}
	dislikes := strings.Split(dislike, " ")
	var dislikesid [][]string
	for _, dislike := range dislikes {
		dislikesid = append(dislikesid, strings.Split(dislike, "="))
	}
	for _, value := range dislikesid {
		if value[0] == "post:"+id_post {
			if value[1] == "-1" {
				value[1] = "0"
				dislike = ""
				for _, nval := range dislikesid {
					if nval[0] != "" {
						dislike += nval[0] + "=" + nval[1] + " "
					}
				}
				_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, dislike, id_user)
				return false
			} else if value[1] == "1" {
				_, _ = db.Exec(`UPDATE postes SET score = score - 1 WHERE id=?;`, id_post)
			}
			value[1] = "-1"
			dislike = ""
			for _, nval := range dislikesid {
				if nval[0] != "" {
					dislike += nval[0] + "=" + nval[1] + " "
				}
			}
			_, _ = db.Exec(`UPDATE user SET like = ? WHERE id = ?;`, dislike, id_user)
			return true
		}
	}
	dislike += " post:" + id_post + "=-1"
	_, _ = db.Exec(`UPDATE user SET dislike = ? WHERE id = ?;`, dislike, id_user)
	return true
}

func SortByLike(val []Value) []Value {
	res := make([]Value, len(val))
	copy(res, val)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})
	return res
}

func SortByOld(val []Value) []Value {
	res := make([]Value, len(val))
	copy(res, val)
	sort.Slice(res, func(i, j int) bool {
		return res[i].Publication_d > res[j].Publication_d
	})
	return res
}
