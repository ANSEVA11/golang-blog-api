//------ In The Name Of God

package main

import (
	"blog-api/config"
	"blog-api/models"
	"blog-api/routes"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB() // جهت اتصال به دیتابیس

	// config.DB.Migrator().DropTable(&models.User{}) DELETE recodes of table
	
	config.DB.AutoMigrate(&models.User{}, &models.Article{}) // ساختن جدول کاربر در دیتابیس براساس مدل کاربر خودمون

	// وصل کردن روتر به توابع هندلر
	router := routes.RegisterRoutes()

	
	fmt.Println("Start Web Server...")
	http.ListenAndServe(":8000", router)

}

