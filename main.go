package main

import (
	"user-service/database"
	"user-service/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()

	r.Run(":8080")
}
