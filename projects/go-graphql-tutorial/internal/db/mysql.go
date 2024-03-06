package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // init mysql driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // init migrate source files
)

var Db *sql.DB

func Init() {
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func Close() error {
	return Db.Close()
}

func Migrate() error {
	if err := Db.Ping(); err != nil {
		return err
	}
	driver, err := mysql.WithInstance(Db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
