package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


func load() {
	err:=godotenv.Load()
	if err != nil{
		fmt.Println( "error while loading the dotenv ,",err)
	}
}

func init(){
	godotenv.Load()
}


func GetString(config_name string) string {
	value,ok := os.LookupEnv(config_name)
	
	if !ok{
		fmt.Println( "Error happend while checking the string in env ")
	}

	return value
}

func GetInt(config_name string) int{

	value,ok := os.LookupEnv(config_name)

	if !ok{
		fmt.Println("Error while checking int in env")
	}
	intValue,err:=strconv.Atoi(value)
	
	if err!= nil{
		fmt.Println("error while converting to integer")
	}

	return intValue
}