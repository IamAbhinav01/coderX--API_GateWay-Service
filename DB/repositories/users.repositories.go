package repositories

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(_name string,email string,password string)(int,error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository( _db *sql.DB)UserRepository{
	return &UserRepositoryImpl{
		db:_db ,
	}
}

func (op *UserRepositoryImpl) Create(_name string,email string,password string) (int,error){

	query := `INSERT INTO users (name,email,password) VALUES (?,?,?)`
	output,err := op.db.Exec(query,_name,email,password)

	if err!= nil{
		fmt.Printf("Error happend while signup user to database --ERROR FOUND ON REPOSITORY LAYER")
		return 0,err
	}

	response,rowErr:=output.RowsAffected()

	if rowErr!= nil{
		fmt.Printf("Error happend while checking number of rows affected")
		return 0,rowErr
	}

	if response == 0{
		fmt.Printf("No rows where affected , Error occured while checking")
		return 0,nil
	}

	fmt.Printf("Succesfully added user to database")

	return int(response),nil

}