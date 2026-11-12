package test

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "golang_user:golang_qwer@tcp(localhost:3306)/belajar_golang_rest_api_test")
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

func setupRouter() http.Handler {
	validate := validator.New()
	DB := setupTestDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, DB, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategoryTable(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE category")
	helper.PanicIfError(err)
}

func TestCreateCategorySuccess(t *testing.T) {
	DB := setupTestDB()
	truncateCategoryTable(DB)
	router := setupRouter()

	requestBody := strings.NewReader(`{"name": "Test Category 4"}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/categories", requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, float64(201), responseBody["code"].(float64))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Test Category 4", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	DB := setupTestDB()
	truncateCategoryTable(DB)
	router := setupRouter()

	requestBody := strings.NewReader(`{"nae": "Test Category 4"}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/categories", requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, float64(400), responseBody["code"].(float64))
	assert.Equal(t, "Bad Request", responseBody["status"])
}
