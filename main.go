package main

import (
	"flag"
	"fmt"
	"go-api/database"
	"go-api/routes"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error : Failed Load .ENV File")
	}
	appPort := os.Getenv("APP_PORT")

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, errParse := strconv.Atoi(os.Getenv("DB_PORT"))
	if errParse != nil {
		panic("Error : Failed Port Conversion")
	}

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	database.CONFIG, err = gorm.Open(mysql.Open(dbUrl))
	if err != nil {
		panic("Error : Failed To Connect Database")
	}
	fmt.Print(dbName)

	migrate := flag.String("m", "", "Unsupport Command")
	flag.Parse()
	command := *migrate

	if command == "migrate" {
		database.Migrate()
		fmt.Printf("Successfull Migrate")
		return
	}
	database.Migrate()
	fmt.Printf("Successfull Migrate")
	server := routes.InitRoutes()
	port := fmt.Sprintf(":%s", appPort)
	server.Run(port)
}
