package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"fmt"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "golang_user:golang_qwer@tcp(127.0.0.1:3306)/belajar_golang_rest_api")
	helper.PanicIfError(err)
	err = db.Ping()
	helper.PanicIfError(err)
	fmt.Println("Database connected successfully!")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	db.SetConnMaxLifetime(60 * 60)
	db.SetConnMaxIdleTime(10 * 60)
	return db
}
