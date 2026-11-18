package app

import (
	"belajar-golang-restful-api/config"
	"belajar-golang-restful-api/helper"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewDB(cfg *config.Config) *sql.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Println("Database connected successfully!")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
