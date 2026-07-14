package services

import (
	"Coderx/DB/repositories"
	"Coderx/dtos"
	"Coderx/security"
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

	config,configErr := security.NewArgonConfig()

	if configErr != nil{
		fmt.Printf("Error occured while configuring the security parameters")
		return 0,configErr
	}

	hashedPassword , hashErr := config.HashPassword(payload.Password)

	if hashErr != nil{
		fmt.Printf("Error Occured while hashing the password")
		return 0,hashErr
	}


	response , err:= user.repo.Create(payload.Name,payload.Email,hashedPassword)

	if err != nil{
		fmt.Println("Error while forwarding request from service to repository")
		return 0,err
	}
	return response,nil
}

