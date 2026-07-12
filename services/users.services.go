package services

import (
	"Coderx/DB/repositories"
	"Coderx/dtos"
	"fmt"
)

type UserService interface {
	SignUp(payload dtos.SignupRequestDTO)(int,error)
}

type UserServiceImp struct {
	repo repositories.UserRepository
}

func NewService(_repo repositories.UserRepository) UserService{
	return &UserServiceImp{
		repo: _repo,
	}
}

func (user *UserServiceImp) SignUp(payload dtos.SignupRequestDTO) (int,error){

	response , err:= user.repo.Create(payload.Name,payload.Email,payload.Password)

	if err != nil{
		fmt.Println("Error while forwarding request from service to repository")
		return 0,err
	}
	return response,nil
}

