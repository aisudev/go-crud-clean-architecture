package main

import (
	"crud/domain"
	"crud/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	martDelivery "crud/feature/mart/delivery"
	martRepository "crud/feature/mart/repository"
	martUsecase "crud/feature/mart/usecase"
)

var postgresDB *gorm.DB

func init() {

	utils.InitViper()

	var err error
	postgresDB, err = newPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate()
}

func main() {
	e := echo.New()

	group := e.Group("")

	martDelivery.NewHandler(
		group,
		martUsecase.NewMartUsecase(
			martRepository.NewMartRepository(postgresDB),
		),
	)

	e.Logger.Fatal(e.Start(utils.ViperGetString("http.port")))
}

func newPostgresDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		utils.ViperGetString("postgres.host"),
		utils.ViperGetString("postgres.user"),
		utils.ViperGetString("postgres.password"),
		utils.ViperGetString("postgres.dbname"),
		utils.ViperGetString("postgres.port"))

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return db, err
}

func AutoMigrate() {
	err := postgresDB.AutoMigrate(
		&domain.Mart{},
	)

	if err != nil {
		panic(err)
	}
}
