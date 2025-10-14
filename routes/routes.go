//------ In The Name Of God

package routes

import (
	"blog-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	//------ User Routes -------
	router.HandleFunc("/api/v1/create_user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/get_users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/get_user/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/update_user/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("api/v1/delete_user/{id}", handlers.DeleteUser).Methods("DELETE")

	//------ Article Routes -------
	router.HandleFunc("/api/v1/create_article", handlers.CreateArticle).Methods("POST")
	router.HandleFunc("/api/v1/get_articles", handlers.GetArticles).Methods("GET")
	router.HandleFunc("/api/v1/get_article/{id}", handlers.GetArticle).Methods("GET")
	router.HandleFunc("/api/v1/update_article/{id}", handlers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/api/v1/delete_article/{id}", handlers.DeleteArticle).Methods("DELETE")

	return router

}