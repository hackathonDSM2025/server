package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"hackathon-dsm-server/internal/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewConnection() *sql.DB {
	cfg := config.GlobalConfig.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	runMigrations(db)

	log.Println("Successfully connected to MySQL")
	return db
}

func runMigrations(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}