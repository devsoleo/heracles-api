package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import pour enregistrer le driver sqlite
)

var db *sql.DB

func InitDB() {
	var err error

	// Ouvre (ou crée) db.sqlite à la racine du projet
	db, err = sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatalf("Erreur d'ouverture de la base de données: %v", err)
	}

	// Test de connexion
	if err = db.Ping(); err != nil {
		log.Fatalf("Erreur de connexion à la base de données: %v", err)
	}

	log.Println("Connexion à SQLite établie ✅")
}

func GetDB() *sql.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatalf("Erreur de fermeture de la base de données: %v", err)
	}
}
