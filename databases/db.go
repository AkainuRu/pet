package databases

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Message struct {
	ID   int    `json:"ID"`
	Text string `json:"text"`
}



func InitDB() {
	dsn := "host=localhost user=postgres password=mypassword dbname=postgres port=5430 sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Неудалось подключиться к базе данных: %v", err)
	}
	
	Db.AutoMigrate(&Message{})
	
}
