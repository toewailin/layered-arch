package config

import (
	"app/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB will hold the global DB connection object
var DB *gorm.DB

// LoadConfig initializes the MySQL connection using environment variables
func LoadConfig() error {
	// Load environment variables from .env file
	err := godotenv.Load("./config/.env")  // Specify path explicitly
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return err
	}

	// Get database credentials from environment variables
	dbHost := GetEnv("DB_HOST")
	dbPort := GetEnv("DB_PORT")
	dbUser := GetEnv("DB_USER")
	dbPassword := GetEnv("DB_PASSWORD")
	dbName := GetEnv("DB_NAME")

	// Build connection string for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the MySQL database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to MySQL database:", err)
		return err
	}

	// Automatically migrate the schema (create the table if it doesn't exist)
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Could not migrate database schema:", err)
		return err
	}

	log.Println("Successfully connected to the database")
	return nil
}

// GetEnv retrieves environment variable value for a given key
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set!", key)
	}
	return value
}
