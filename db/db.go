package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(connectionString string) (Database, error) {
	db := Database{}
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	errMigrate := migrateDB(db)

	if errMigrate != nil {
		return db, errMigrate
	}
	return db, nil
}

func migrateDB(db Database) error {
	driver, err := postgres.WithInstance(db.Conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)

	if err != nil {
		log.Println(err)
		return err
	}
	m.Steps(2)
	log.Println("Database migrations complete")
	return nil
}
