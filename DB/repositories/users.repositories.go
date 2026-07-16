package repositories

import (
	"Coderx/schema"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create(_name string,email string,password string)(int,error)
	GetCredentialByEmail(email string) (*schema.User,error)
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

	_,rowErr:=output.RowsAffected()
	response,inserErr := output.LastInsertId()

	if inserErr!= nil{
		fmt.Println("Erro happend while returning user id from db")
		return 0,inserErr
	}


	if rowErr!= nil{
		fmt.Printf("Error happend while checking number of rows affected")
		return 0,rowErr
	}

	fmt.Printf("Succesfully added user to database")

	return int(response),nil

}

func (op *UserRepositoryImpl) GetCredentialByEmail(email string) (*schema.User,error){

	user := schema.User{}

	query:= `select id, password from users where email = ?`

	err := op.db.QueryRow(query,email).Scan(&user.ID,&user.Password)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, fmt.Errorf("user not found")
		}
		fmt.Printf("Error fetching user by email: %v\n", err)
		return nil, err
	}
	return &user,nil

}