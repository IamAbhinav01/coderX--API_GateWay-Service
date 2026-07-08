package db

import (
	"Coderx/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DBUSER")
	cfg.Passwd = env.GetString("DBPASS")
	cfg.Net = env.GetString("DB_Net")
	cfg.Addr = env.GetString("DB_Addr")
	cfg.DBName = env.GetString("DBName")


	DB,err:=sql.Open("mysql",cfg.FormatDSN())

	if err != nil{
		fmt.Println("Error in initialsing database")
	}

	pingErr:=DB.Ping()

	if pingErr != nil{
		fmt.Print("Error connecting to database")
	}

	fmt.Println("Connected to database")

	return DB,nil
}