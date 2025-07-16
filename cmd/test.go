package main

import (
	"fmt"

	"github.com/Akiles94/go-test-api/config"
	"github.com/Akiles94/go-test-api/db"
)

func main() {
	config.LoadEnv()
	db := db.Connect()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("✅ DB connection OK")
}
