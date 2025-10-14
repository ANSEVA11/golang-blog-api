//------ In The Name Of God

package handlers

import (
	"blog-api/config"
	"blog-api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//-----Create Article------
func CreateArticle(w http.ResponseWriter, r *http.Request){
	var article models.Article
	json.NewDecoder(r.Body).Decode(&article)

	result := config.DB.Create(&article)
	if result.Error !=nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Preload("User").First(&article, article.ID)

	json.NewEncoder(w).Encode(article)
}

//-----Get Articles------
func GetArticles(w http.ResponseWriter, r *http.Request){
	var articles []models.Article
	config.DB.Preload("User").Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

//-----Get Article------
func GetArticle(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)

	id_str := parameters["id"]
	id, _ := strconv.Atoi(id_str)

	var article models.Article

	result := config.DB.Preload("User").First(&article, id)
	if result.Error != nil {
		http.Error(w, "Article Not Found!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

//-----Update Article------
func UpdateArticle(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)

	id_str := parameters["id"]
	id, _ := strconv.Atoi(id_str)

	var article models.Article
	result := config.DB.First(&article, id)
	if result.Error != nil {
		http.Error(w, "Article Not Found!", http.StatusBadRequest)
		return
	}

	var new_article models.Article
	json.NewDecoder(r.Body).Decode(&new_article)

	article.Title = new_article.Title
	article.Body = new_article.Body

	result = config.DB.Save(&article)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Preload("User").First(&article, id)
	json.NewEncoder(w).Encode(article)
}

//-----Delete Article------
func DeleteArticle(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	id_str := parameters["id"]
	id, _ := strconv.Atoi(id_str)

	var article models.Article
	result := config.DB.First(&article, id)
	if result.Error != nil {
		http.Error(w, "Article Not Found!", http.StatusNotFound)
		return
	}

	config.DB.Delete(&article)
	w.WriteHeader(http.StatusNoContent)

}

