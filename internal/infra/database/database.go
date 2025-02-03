package database

import (
	"database/sql"
	"fmt"
	"go-kafka-order-producer/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect() *sql.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBName,
	)

	db, err := sql.Open(config.Config.DBDriver, connectionString)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		panic(err)
	}

	return db
}

func Migrate(db *sql.DB) {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://../internal/database/migrations",
		config.Config.DBDriver,
		driver,
	)
	if err != nil {
		panic(err)
	}

	m.Up()
}
