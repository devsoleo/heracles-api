package main

import (
	"devsoleo/heracles-back/database"
	"devsoleo/heracles-back/router"
)

func main() {
	database.InitDB()
	defer database.Close()

	r := router.SetupRouter()
	r.Run(":8080")
}
