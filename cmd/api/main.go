package main

import (
	"cashflow/internal/controllers"
	"cashflow/internal/domain/transaction"
	"cashflow/internal/infraestructure/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(logger.New())

	transactionService := transaction.TransactionService{
		Repository: &database.TransactionRepository{},
	}
	controller := controllers.ControllerBase{
		TransactionService: transactionService,
	}

	app.Post("/transactions", controller.CreateTransaction)
	app.Get("/transactions", controller.GetTransactions)

	log.Fatal(app.Listen(":3000"))
}
