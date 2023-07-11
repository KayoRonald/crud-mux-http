package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gtihub.com/KayoRonald/crud-mux-http/models"
)

type DbInstance struct {
  Db *gorm.DB
}

var Database DbInstance

func ConnectDB(){
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
  if err != nil {
    log.Panic("Failed to connect to the database")
    os.Exit(2)
  }
  log.Print("Connect database sucess")

  log.Print("Running Migrations")
  db.AutoMigrate(&models.Tasks{})
  Database = DbInstance{
    Db: db,
  }
}