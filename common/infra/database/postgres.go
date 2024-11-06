package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	maxOpenDbConn = 25              // Número máximo de conexões abertas ao banco de dados
	maxIdleDBConn = 25              // Número máximo de conexões inativas no pool de conexões
	maxDBLifetime = 5 * time.Minute // Tempo máximo de vida de uma conexão no pool
)

type Postgres struct{}

func (d Postgres) InitDB(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return
}

func (d Postgres) RunMigrate(dsn string) (err error) {
	migrationPath := "migrations"

	m, err := migrate.New(
		fmt.Sprintf("file:%s", migrationPath),
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
