package database

import (
	"fmt"
	"glyph/models"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (skipping)")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return value
}

func InitializeDatabase() *gorm.DB {
	host := GetEnvVariable("POSTGRES_HOST")
	username := GetEnvVariable("POSTGRES_USER")
	password := GetEnvVariable("POSTGRES_PASSWORD")
	dbName := GetEnvVariable("POSTGRES_DB")
	port := GetEnvVariable("POSTGRES_PORT")

	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Couldn't convert port number")
	}

	dbConn, err := New(Config{
		Host:     host,
		Port:     portNum,
		User:     username,
		Password: password,
		DBName:   dbName,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := Migrate(dbConn); err != nil {
		log.Fatal(err)
	}

	return dbConn
}

func New(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(log.Writer(), "[gorm] ", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Article{},
		&models.Comment{},
	)
}
