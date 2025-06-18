package main

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        dbUser, dbPass, dbHost, dbPort, dbName)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = db
    log.Println("Database connected successfully")
}
