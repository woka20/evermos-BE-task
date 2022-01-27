package databases

import (
	"errors"
	"evermos-be-task/task-store/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	migrate "github.com/rubenv/sql-migrate"
)

type Database struct {
	DB   *gorm.DB
	Init bool
}

type DatabaseInterface interface {
	DBInit() error
	Migrate() error
}

type DatabaseRepo struct {
}

func NewDatabaseRepo() DatabaseInterface {
	return DatabaseRepo{}
}

var database Database

func (d DatabaseRepo) DBInit() error {

	DbUser, DbPassword, DbPort, DbHost, DbName := config.DB_UNAME, config.DB_PASS, config.DB_PORT, config.DB_HOST, config.DB_NAME
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	database = Database{DB: db, Init: true}

	return nil

}

func (d DatabaseRepo) Migrate() error {
	migrations := &migrate.FileMigrationSource{Dir: "migrations"}

	if database.Init == true {
		dbConn := database.DB
		fmt.Println("Starting migration")
		result, err := migrate.Exec(dbConn.DB(), "mysql", migrations, migrate.Up)

		if err != nil {
			log.Fatalf("Migration failed: %s", err)
		}

		log.Printf("Applied %d migrations successfully", result)

		return nil

	} else {

		return errors.New("DB Not Initialized")
	}

}
