package main

import (
	"assigment-2/database"
	"assigment-2/router"
	"fmt"
)

func main() {
	err := database.PostGresDB()

	if err != nil {
		fmt.Println("Database isn't running right now", err)
	}

	router.Routers().Run()
}
