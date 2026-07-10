package services

import (
	"Coderx/DB/repositories"
	"fmt"
)

type UserService interface {
	SignUp(_name string, email string, password string)(int,error)
}

type UserServiceImp struct {
	repo repositories.UserRepository
}

func NewService(_repo repositories.UserRepository) UserService{
	return &UserServiceImp{
		repo: _repo,
	}
}

func (user *UserServiceImp) SignUp(_name string, email string, password string) (int,error){

	response , err:= user.repo.Create(_name,email,password)

	if err != nil{
		fmt.Println("Error while forwarding request from service to repository")
		return 0,err
	}
	return response,nil
}

