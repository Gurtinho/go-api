package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/gurtinho/go/api/internal/dtos"
	"github.com/gurtinho/go/api/internal/entities"
	"github.com/gurtinho/go/api/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Authenticate User Godoc
// @Summary Authenticate User
// @Description Authenticate User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dtos.GetJWTDTO true "user credentials"
// @Success 200 {object} dtos.GetJwtOutput
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /users/generate-token [post]
func (usr *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("expiresIn").(int)
	var user dtos.GetJWTDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := usr.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
	})
	accessToken := dtos.GetJwtOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create User Godoc
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dtos.UserDTO true "user request"
// @Success 201
// @Failure 500 {object} Error
// @Router /users [post]
func (usr *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entities.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = usr.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
