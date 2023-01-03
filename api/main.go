package main

import (
	"log"

	"example.com/mod/delivery"
	"example.com/mod/repository"
	"example.com/mod/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	tr := repository.NewSyncMapTodoRepository()
	tu := usecase.NewTodoUsecase(tr)
	c := fiber.New()

	c.Use(cors.New(cors.Config {
		AllowCredentials: true,
	}))

	delivery.NewTodoAllGetHandler(c, tu)
	delivery.NewTodoDeleteHandler(c, tu)
	delivery.NewTodoStatusUpdateHandler(c, tu)
	delivery.NewTodoStoreHandler(c, tu)

	log.Fatal(c.Listen("localhost:8080"))
}
