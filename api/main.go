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

	// fiber >= 2.52 rejects AllowCredentials with a wildcard origin,
	// so list the SvelteKit dev (5173) and preview (4173) origins explicitly.
	c.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:4173",
		AllowCredentials: true,
	}))

	delivery.NewTodoAllGetHandler(c, tu)
	delivery.NewTodoDeleteHandler(c, tu)
	delivery.NewTodoStatusUpdateHandler(c, tu)
	delivery.NewTodoStoreHandler(c, tu)

	log.Fatal(c.Listen("localhost:8080"))
}
