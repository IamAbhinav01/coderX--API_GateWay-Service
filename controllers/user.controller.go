package controllers

import (
	"Coderx/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewController(_user_service services.UserService) *UserController{
	return &UserController{
		UserService: _user_service,
	}
}

func (controller *UserController) SignUp(w http.ResponseWriter, r *http.Request) {

	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing form data"))
		return
	}

	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	response, err := controller.UserService.SignUp(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully. Rows affected: %d", response)

}

