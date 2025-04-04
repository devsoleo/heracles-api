package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type User struct {
	Entry  int    `json:"entry"`
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./src/database.db")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données : " + err.Error())
	}
}

func searchUsers(c *gin.Context) {
	query := c.Query("query")
	lang := c.Query("lang")

	if query == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Les paramètres 'query' et 'lang' sont requis"})
		return
	}

	rows, err := db.Query("SELECT * FROM creatures WHERE name LIKE ? AND locale = ?", "%"+query+"%", lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Entry, &user.Name, &user.Locale); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*", "http://localhost"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.GET("/search", searchUsers)

	r.Run(":8080")
}
