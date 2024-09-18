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

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
