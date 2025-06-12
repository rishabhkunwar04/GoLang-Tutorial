
## Controller-Service-Repository (DAO)-DTO-Model pattern
 **Advantage of this pattern**
 1.  Each layer does one thing well
 2. You can mock the repository for unit tests.
 3. Changes in business logic won’t affect controller or DB.




```go
🔁 Overview of Each Layer
Layer	                          Responsibility
Controller:	                 Entry point for requests (e.g., HTTP handlers); handles parsing and response.
DTO	Data Transfer Object:   represents input/output format for APIs.
Service	Business logic:     validates, transforms, coordinates components.
Repository/DAO:	            Handles data persistence; abstracts DB calls (CRUD operations).
Model:	                    Represents the database schema or domain entity.

🧠 Example: User Management in Go (Create User)

🧩 Folder Structure
/user
  - controller.go
  - service.go
  - repository.go
  - model.go
  - dto.go
  - handler.go

1. models/user.go

package models

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
2. dto/user_dto.go

package dto

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
3. repository/user_repository.go

package repository

import (
	"example.com/project/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByID(id int64) (models.User, error)
}
Implementation:

package repository

import (
	"database/sql"
	"example.com/project/models"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user models.User) (models.User, error) {
	res := r.db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email)
	err := res.Scan(&user.ID)
	return user, err
}

func (r *userRepositoryImpl) FindByID(id int64) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}
4. service/user_service.go

package service

import (
	"example.com/project/dto"
	"example.com/project/models"
	"example.com/project/repository"
)

type UserService interface {
	CreateUser(dto.CreateUserRequest) (dto.UserResponse, error)
	GetUser(id int64) (dto.UserResponse, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userServiceImpl{repo: r}
}

func (s *userServiceImpl) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	created, err := s.repo.Create(user)
	return dto.UserResponse(created), err
}

func (s *userServiceImpl) GetUser(id int64) (dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	return dto.UserResponse(user), err
}

5. controller/user_handler.go

package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/project/dto"
	"example.com/project/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service service.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	res, err := h.Service.CreateUser(req)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	res, err := h.Service.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

6. main.go

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"

	"example.com/project/controller"
	"example.com/project/repository"
	"example.com/project/service"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=users sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	handler := &controller.UserHandler{Service: svc}

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```