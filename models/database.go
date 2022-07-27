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
func SetupDatabase() *gorm.DB {
	u := helper.GetEnv("DATABASE_USER", "")
	p := helper.GetEnv("DATABASE_PASSWORD", "")
	// h := helper.GetEnv("DATABASE_HOST", "")
	n := helper.GetEnv("DATABASE_NAME", "")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?%s", u, p, n, q)
	fmt.Println(dsn)
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

	return DB
}
