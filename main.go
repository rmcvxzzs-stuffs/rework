package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	db, err := sql.Open("mysql", cfg.MySQL.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	RegisterAnnouncementRoutes(app, db)
    AppName: "Rework",

	// Add other route registrations here (i.e., matchmaking, lobby, etc.)

	log.Fatal(app.Listen(":8080"))
}
