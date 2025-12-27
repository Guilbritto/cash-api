// @title           API de Controle Financeiro
// @version         1.0
// @description     Backend em Go para controle financeiro pessoal.
// @termsOfService  http://example.com/terms/

// @contact.name   Guilherme
// @contact.email  voce@example.com

// @host      localhost:3000
// @BasePath  /api
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Guilbritto/cash-api/internal/application/category"
	"github.com/Guilbritto/cash-api/internal/application/transaction"
	"github.com/Guilbritto/cash-api/internal/controllers"
	_ "github.com/Guilbritto/cash-api/internal/docs"
	"github.com/Guilbritto/cash-api/internal/infraestructure/database"
	"github.com/Guilbritto/cash-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Printf("Error loading .env: %v", err)
	}

	app := fiber.New()

	// Registrando middlewares globais
	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	db := database.NewDb()

	transactionService := transaction.TransactionService{
		TransactionRepository: &database.TransactionRepository{Db: db},
		CategoryRepository:    &database.CategoryRepository{Db: db},
	}
	categoryService := category.CategoryService{
		CategoryRepository: &database.CategoryRepository{Db: db},
	}
	controller := controllers.ControllerBase{
		TransactionService: transactionService,
		CategoryService:    categoryService,
	}

	api := app.Group("/api", middleware.AuthMiddleware)
	api.Post("/transactions", controller.CreateTransaction)
	api.Get("/transactions", controller.GetTransactions)
	api.Post("/category", controller.CreateCategory)

	log.Fatal(app.Listen(":3000"))
}

func loadEnv() error {
	if err := godotenv.Load(); err == nil {
		return nil
	}

	envPath, err := findEnvFile()
	if err != nil {
		return err
	}

	return godotenv.Load(envPath)
}

func findEnvFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
