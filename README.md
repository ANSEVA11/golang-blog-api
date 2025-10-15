# 📰 Blog RESTful API with Golang, GORM, and PostgreSQL
A simple RESTful API project built using Golang, Gorilla Mux, GORM, and PostgreSQL.
This project simulates a simple blog system where users can create, read, update, and delete articles.

---
## 🚀 Features
- Full CRUD operations for Users and Articles
- Connection between Golang and PostgreSQL using GORM ORM
- JSON-based communication
- Error handling
- Organized and scalable folder structure

---
## 🧠 Technologies Used
| Category        | Tools / Libraries       |
| --------------- | ----------------------- |
| Backend         | Golang (1.23+)          |
| Framework       | Gorilla Mux             |
| ORM             | GORM                    |
| Database        | PostgreSQL              |
| Tools           | Postman for API testing |
| Version Control | Git & GitHub            |

---

## 📁 Project Structure
```
blog-api/
├── config/
│   └── db.go        # Database connection
├── handlers/
│   ├── userHandler.go     # CRUD for Users
│   └── articleHandler.go  # CRUD for Articles
├── models/
│   ├── user.go
│   └── article.go
├── routes/
│   └── routes.go
├── main.go
├── go.mod
└── go.sum

```
---

## ⚙️ Setup & Installation
1️⃣ Clone the Repository
```
git clone https://github.com/ANSEVA11/golang-blog-api.git
cd golang-blog-api
```
2️⃣ Install Dependencies
```
go mod tidy
```
3️⃣ Configure PostgreSQL
  Create a database named blog_db (or your custom name).
  - Update your database connection in config/db.go:
  ```
  dsn := "host=localhost user=postgres password=yourpassword dbname=blog_db port=5432 sslmode=disable"
  ```
4️⃣ Run Auto Migration
  Automatically creates required tables ```(users, articles)``` in PostgreSQL.
  ```
  config.DB.AutoMigrate(&models.User{}, &models.Article{})
  ```
5️⃣ Start the Server
```
go run main.go
```
- Server will start at: ```http://localhost:8080```

---
## 🧩 API Endpoints
- 👤 User Routes
  
  | Method | Endpoint      | Description       |
  | ------ | ------------- | ----------------- |
  | GET    | `/users`      | Get all users     |
  | GET    | `/users/{id}` | Get a user by ID  |
  | POST   | `/users`      | Create a new user |
  | PUT    | `/users/{id}` | Update a user     |
  | DELETE | `/users/{id}` | Delete a user     |

- 📝 Article Routes
  
  | Method | Endpoint         | Description          |
  | ------ | ---------------- | -------------------- |
  | GET    | `/articles`      | Get all articles     |
  | GET    | `/articles/{id}` | Get an article by ID |
  | POST   | `/articles`      | Create a new article |
  | PUT    | `/articles/{id}` | Update an article    |
  | DELETE | `/articles/{id}` | Delete an article    |


