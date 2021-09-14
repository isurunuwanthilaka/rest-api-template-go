package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB //database
)

func init() {

	e := godotenv.Load("dev.env") //Load .env file
	if e != nil {
		log.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbName) //Build connection string

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}

	db = conn
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
