package controllers

import (
	"Coderx/dtos"
	"Coderx/services"
	"Coderx/utils/formatters"

	"net/http"
	"strings"
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

	payload := r.Context().Value("payload").(dtos.SignupRequestDTO)

	response, err := controller.UserService.SignUp(payload)

	if err != nil {
		status:=http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()),"duplicate") || strings.Contains(err.Error(), "1062"){
			status = http.StatusConflict
		}
		formatters.ErrorResponse(w,status,"Error occured while signing the user",err)
		return
	}
	formatters.SuccessResponse(w,http.StatusCreated,"User sign-up successfully",response)


}

