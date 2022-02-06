package models

import (
	"fmt"

	"github.com/vincentJunior1/test-kriya/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// Connect to the database.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Migrate the schema

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}
