package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error while loading the dotenv,", err)
	}
}

func init() {
	godotenv.Load()
}

func GetString(config_name string) string {
	godotenv.Load()
	value, ok := os.LookupEnv(config_name)

	if !ok {
		return "env: missing required key"
	}

	return value
}

func GetInt(config_name string) int {
	godotenv.Load()
	value, ok := os.LookupEnv(config_name)

	if !ok {
		return 0
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return intValue
}
