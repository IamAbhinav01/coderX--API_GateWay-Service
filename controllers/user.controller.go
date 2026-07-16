package controllers

import (
	"Coderx/dtos"
	"Coderx/middlewares"
	"Coderx/services"
	"Coderx/utils/formatters"
	"Coderx/utils/session"
	"fmt"
	"strconv"

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

	userSession := &session.Session{
		Data: map[string]string{
			"user_id": strconv.Itoa(response),
			"email":payload.Email,
		},
	}

	err = controller.SessionManager.Migrate(r.Context(),userSession)
	if err == nil{
		cookie:= &http.Cookie{
			Name: "session_id",
			Value: userSession.Id,
			Path: "/",
			MaxAge: 86400,
			HttpOnly: true,
			Secure: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w,cookie)
	}

	formatters.SuccessResponse(w,http.StatusCreated,"User sign-up successfully",response)


}

func (contoller *UserController) Login(w http.ResponseWriter,r *http.Request){


	payload := r.Context().Value(middlewares.PayloadContextKey).(dtos.LoginRequestDTO)
	fmt.Println(payload)

}
