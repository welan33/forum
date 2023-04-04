package main

import (
	f "forum/fonctions"
	"log"
	"net/http"
)

func main() {
	forum_db := f.InitDB("forum.db")
	f.CreateTable(forum_db)
	defer forum_db.Close()
	println("Récupération de la base de donnée effectuée")

	f.HandleRequests()
	//f.AlterDB()

	//retirer le /static de l'url
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//retirer le /assets de l'url
	fsAssets := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))

	//run serveur
	println("Listening at http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
