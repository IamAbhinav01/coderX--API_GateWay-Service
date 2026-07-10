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

func (controller *UserController) SignUp(w http.ResponseWriter , r http.Request){

	response,err:=controller.UserService.SignUp("Abhinav","AbhinavSunil70@gmail.com","its785Abbhina")
	
	fmt.Println("reponse is : ",response)
	if err != nil{
		w.Write([]byte ("user request succesfully sent to service layer from controller"))
	}

}

