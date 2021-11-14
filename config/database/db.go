package database

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"quiz-app/models"
	"strconv"
)

var DB *gorm.DB

func Connect() error {
	// load the .env file
	_ = godotenv.Load()

	var server = os.Getenv("SERVER")
	var port, _ = strconv.Atoi(os.Getenv("PORT_DB"))
	var user = os.Getenv("USER")
	var password = os.Getenv("PASSWORD")
	var database = os.Getenv("DATABASE")

	// Build connection string
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	fmt.Println(dsn)
	// Create connection pool
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error creating connection pool: %s\n", err.Error())
		return err
	}

	// auto migrate
	if err = DB.AutoMigrate(&models.CourseQuiz{}, &models.QuestionQuiz{}); err != nil {
		return err
	}

	fmt.Printf("Connected!\n")
	return nil
}
