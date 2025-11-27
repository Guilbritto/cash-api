package main

import (
	"log"

	"github.com/Guilbritto/cash-api/internal/controllers"
	"github.com/Guilbritto/cash-api/internal/domain/transaction"
	"github.com/Guilbritto/cash-api/internal/infraestructure/database"
	"github.com/Guilbritto/cash-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(middleware.AuthMiddleware)

	db := database.NewDb()

	transactionService := transaction.TransactionService{
		Repository: &database.TransactionRepository{Db: db},
	}
	controller := controllers.ControllerBase{
		TransactionService: transactionService,
	}

	app.Post("/transactions", controller.CreateTransaction)
	app.Get("/transactions", controller.GetTransactions)

	log.Fatal(app.Listen(":3000"))
}
