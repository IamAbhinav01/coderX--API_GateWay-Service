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

	if cfg.User == "" {
		return nil, fmt.Errorf("missing required env DBUSER")
	}
	if cfg.Passwd == "" {
		return nil, fmt.Errorf("missing required env DBPASS")
	}
	if cfg.Net == "" {
		return nil, fmt.Errorf("missing required env DB_Net")
	}
	if cfg.Addr == "" {
		return nil, fmt.Errorf("missing required env DB_Addr")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("missing required env DBName")
	}

	DB, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		fmt.Println("Error connecting to database:", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to database successfully")

	return DB, nil
}