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


	response, err := controller.UserService.SignUp("Abhinav", "abhinavsunil70@gmail.com", "abhinavs784d")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully. Rows affected: %d", response)

}

