package db

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"tunwa/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Log struct {
	ID        uint      `gorm:"primary_key"`
	Request   string    `gorm:"type:jsonb"`
	Response  string    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

func InitDB() {
	// Load configuration
	config.LoadConfig()

	// Construct the Data Source Name (DSN)
	dsn := "host=" + config.AppConfig.DBHost +
		" user=" + config.AppConfig.DBUser +
		" password=" + config.AppConfig.DBPassword +
		" dbname=" + config.AppConfig.DBName +
		" port=" + config.AppConfig.DBPort +
		" sslmode=" + config.AppConfig.DBSSLMode +
		" TimeZone=" + config.AppConfig.DBTimezone

	// Open the database connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// AutoMigrate database schema
	if err := db.AutoMigrate(&Log{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func LogRequestResponse(request *http.Request, response interface{}) {
	// Convert request to JSON
	reqBody, err := json.Marshal(map[string]interface{}{
		"method":  request.Method,
		"url":     request.URL.String(),
		"headers": request.Header,
	})
	if err != nil {
		log.Printf("Error marshalling request: %v", err)
		return
	}

	// Convert response to JSON
	resBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return
	}

	// Create a new log entry
	logEntry := Log{
		Request:  string(reqBody),
		Response: string(resBody),
	}

	// Save the log entry to the database
	db.Create(&logEntry)
}
