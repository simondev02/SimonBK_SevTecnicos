package main

import (
	"SimonBK_SevTecnicos/domain/models"
	"SimonBK_SevTecnicos/infra/db"
	"fmt"
	"log"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err = db.DBConn.AutoMigrate(&models.TechnicalServiceToStatus{})
	if err != nil {
		log.Fatalf("Failed to migrate table Vehicle: %v", err)
	}

	fmt.Println("Migration successful")
}
