package delivery

import (
	"example.com/mod/domain"

	"github.com/gofiber/fiber/v2"
)

type todoStatusUpdateHandler struct {
	todoUseCase domain.TodoUsecase
}

func NewTodoStatusUpdateHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoStatusUpdateHandler{
		todoUseCase: th,
	}

	c.Post("/todo/update", handler.StatusUpdate)
}

func (h *todoStatusUpdateHandler) StatusUpdate(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	err := c.BodyParser(todo)

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unexpected Request. To check your ID",
		})
	}

	err = h.todoUseCase.StatusUpdate(todo.ID)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON("StatusUpdate OK")
}
