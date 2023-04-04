package forum

import "net/http"

func HandleRequests() {
	//gestion des routes
	http.HandleFunc("/", MainHandler)

	//routes pour authentification
	http.HandleFunc("/register", RegHandler)
	http.HandleFunc("/login", LogHandler)
	http.HandleFunc("/googleRegister", GoogleHandler)
	http.HandleFunc("/googleLogin", GoogleLogHandler)
	http.HandleFunc("/co", CoHandler)
	http.HandleFunc("/logout", OutHandler)
	http.HandleFunc("/profil/", ProfilHandler)
	http.HandleFunc("/admin", UserToAdmin)
	http.HandleFunc("/mod", UserToModerator)
	http.HandleFunc("/vip", UserToVIP)
	http.HandleFunc("/vip-p", VIPHandler)
	http.HandleFunc("/reset", ResetRank)
	http.HandleFunc("/admin-p/", AdminHandler)
	http.HandleFunc("/delete-u/", DeleteuserDB)
	http.HandleFunc("/category", UpdateCategories)
	http.HandleFunc("/delete-c/", DeletecatDB)
	http.HandleFunc("/request", RequestHandler)
	http.HandleFunc("/sendrequest", SaveRequest)
	http.HandleFunc("/delete-r/", DeleterequestDB)

	//routes pour les postes utilisateurs
	http.HandleFunc("/poste", PostHandler)
	http.HandleFunc("/saveposte", InsertPost)
	http.HandleFunc("/comment", CommentHandler)
	http.HandleFunc("/savecomment", InsertComment)
	http.HandleFunc("/EditProfil", ModifProfile)
	http.HandleFunc("/page/", PageDisplay)
	http.HandleFunc("/delete/", DeletepostDB)
	http.HandleFunc("/like/", LikeHandler)
	http.HandleFunc("/dislike/", DislikeHandler)
	http.HandleFunc("/likeC/", LikeCHandler)
	http.HandleFunc("/dislikeC/", DislikeCHandler)
}
