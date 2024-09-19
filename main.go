package main

import (
	"fmt"
	"os"
	"user-service/database"
	"user-service/router"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	database.StartDB()

	r := router.StartApp()

	fmt.Println("User Service running on port:", os.Getenv("PORT"))
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
