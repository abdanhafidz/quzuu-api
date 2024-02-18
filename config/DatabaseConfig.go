package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"github.com/quzuu-be/models"
)

func AutoMigrateAll(db *gorm.DB) {
	// Enable logger to see SQL logs
	db.Logger.LogMode(logger.Info)

	// Auto-migrate all models
	err := db.AutoMigrate(
		&models.ProblemSetAssign{},
		&models.Announcement{},
		&models.Account{},
		&models.MCQuestion{},
		&models.ShortAnsQuestion{},
		&models.EssayQuestion{},
		&models.ProblemSet{},
		&models.AccountDetails{},
		&models.Events{},
		&models.EventAssign{},
		&models.Result{},
		&models.ExamProgress{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully.")
}

var DB *gorm.DB
var err error
var Salt string

func init() {
	godotenv.Load()
	if err != nil {
		fmt.Println("Gagal membaca file .env")
		return
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	Salt := os.Getenv("SALT")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
	if Salt == "" {
		Salt = "D3f4u|t"
	}

	// Call AutoMigrateAll to perform auto-migration
	AutoMigrateAll(DB)
}
