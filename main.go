package main

import (
	"devsoleo/heracles-api/database"
	"devsoleo/heracles-api/router"
)

func main() {
	database.InitDB()
	defer database.Close()

	r := router.SetupRouter()
	r.Run(":8080")
}
