package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vincentJunior1/test-kriya/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	u := helper.GetEnv("DATABASE_USER", "")
	p := helper.GetEnv("DATABSE_PASSWORD", "")
	h := helper.GetEnv("DATABASE_HOST", "")
	n := helper.GetEnv("DATABASE_NAME", "")
	q := "charset=utf8mb4&parseTime=True&loc=Local"
	// :test@tcp(13.228.73.161:8976)/kriya_test
	// Assemble the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	// Connect to the database.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// Migrate the schema

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
