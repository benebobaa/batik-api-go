package database

import (
	"fmt"
	"log"
	"os"

	"github.com/benebobaa/batik-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Could not connect to the database \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating the schema...")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}
