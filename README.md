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

