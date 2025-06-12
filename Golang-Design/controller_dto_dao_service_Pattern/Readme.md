
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
main.go
🧱 model.go — Domain/DB Layer

package user

type User struct {
	ID    int
	Name  string
	Email string
}
📦 dto.go — Data Transfer Object

package user

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
📂 repository.go — DAO Layer

package user

type UserRepository interface {
	Save(user User) (User, error)
}

type InMemoryUserRepo struct {
	users []User
	autoID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{}
}

func (r *InMemoryUserRepo) Save(user User) (User, error) {
	r.autoID++
	user.ID = r.autoID
	r.users = append(r.users, user)
	return user, nil
}
🧠 service.go — Business Logic

package user

type UserService interface {
	CreateUser(req CreateUserRequest) (UserResponse, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req CreateUserRequest) (UserResponse, error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}
	savedUser, err := s.repo.Save(user)
	if err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
		ID:    savedUser.ID,
		Name:  savedUser.Name,
		Email: savedUser.Email,
	}, nil
}
📲 controller.go — HTTP Handler (acts as Controller)

package user

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	userResp, err := uc.service.CreateUser(req)
	if err != nil {
		http.Error(w, "could not create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userResp)
}
🚀 main.go — Entry Point

package main

import (
	"net/http"
	"user"
)

func main() {
	repo := user.NewInMemoryUserRepo()
	service := user.NewUserService(repo)
	controller := user.NewUserController(service)

	http.HandleFunc("/users", controller.CreateUserHandler)
	http.ListenAndServe(":8080", nil)
}
```