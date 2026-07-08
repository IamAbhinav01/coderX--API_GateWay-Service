package main

import (
	"Coderx/app"
	"fmt"
)

func main() {
	application := app.NewApplication()

	if err := application.Run(); err != nil {
		fmt.Println("Error while initialising the server:", err)
	}
}