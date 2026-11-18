package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/config"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize validator
	validate := validator.New()

	// Initialize database
	db := app.NewDB(cfg)
	defer db.Close()

	// Initialize repository
	categoryRepository := repository.NewCategoryRepository()

	// Initialize service
	categoryService := service.NewCategoryService(
		categoryRepository,
		db,
		validate,
	)

	// Initialize controller
	categoryController := controller.NewCategoryController(categoryService)

	// Initialize router
	router := app.NewRouter(categoryController)

	// Initialize middleware
	handler := middleware.CorsMiddleware(
		middleware.NewAuthMiddleware(router, cfg),
		cfg,
	)

	// Initialize server
	server := http.Server{
		Addr:    cfg.AppHost + ":" + cfg.AppPort,
		Handler: handler,
	}

	log.Println("Server running on", server.Addr)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		helper.PanicIfError(err)
	}
}
