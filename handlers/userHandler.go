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

// لیست همه کاربران
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// گرفتن کاربر خاص
func GetUser(w http.ResponseWriter, r *http.Request){
	// گرفتن پارامتر|پارامترها از URL 
	parameters := mux.Vars(r)
	
	id_str := parameters["id"]

	id, _ := strconv.Atoi(id_str)

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// ساخت کاربر جدید
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// ذخیره و چک کردن خطا جهت اینکه اگر ایمیل یا هر چیزی دچار خطا شد
	result := config.DB.Create(&user)
	if result.Error != nil {
    	http.Error(w, result.Error.Error(), http.StatusBadRequest)
    	return
	}
	
	json.NewEncoder(w).Encode(user)
}

// Update Information of User
func UpdateUser(w http.ResponseWriter, r *http.Request){
	// گرفتن پارامتر|پارامترها از URL 
	parameters := mux.Vars(r)

	id_str := parameters["id"]

	id, _ := strconv.Atoi(id_str)

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var new_user models.User
	json.NewDecoder(r.Body).Decode(&new_user)

	user.FullName = new_user.FullName
	user.Email = new_user.Email

	// ذخیره و چک کردن خطا جهت اینکه اگر ایمیل یا هر چیزی دچار خطا شد
	result = config.DB.Save(&user)
	if result.Error != nil {
    	http.Error(w, result.Error.Error(), http.StatusBadRequest)
    	return
	}
	json.NewEncoder(w).Encode(user)

}

//Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)

	id_str := parameters["id"]

	id, _ := strconv.Atoi(id_str)

	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	config.DB.Delete(&user)

	w.WriteHeader(http.StatusNoContent)

}
