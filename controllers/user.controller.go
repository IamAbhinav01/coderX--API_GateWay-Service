package controllers

import (
	"Coderx/dtos"
	"Coderx/middlewares"
	"Coderx/services"
	"Coderx/utils/formatters"
	"Coderx/utils/session"

	"net/http"
	"strings"
)

type UserController struct {
	UserService services.UserService
	SessionManager *session.SessionManager
}

func NewController(_user_service services.UserService, _sm *session.SessionManager) *UserController{
	return &UserController{
		UserService: _user_service,
		SessionManager: _sm,
	}
}

func (controller *UserController) SignUp(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value(middlewares.PayloadContextKey).(dtos.SignupRequestDTO)

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

