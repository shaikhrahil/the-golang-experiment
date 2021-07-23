package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

type controller struct {
	logger         *log.Logger
	accountService accounts.Repository
	authService    auth.Repository
	todoService    Repository
}

func NewController(r *fiber.Router, logger *log.Logger, authService auth.Repository, accountsService accounts.Repository) {
	router := *r
	h := controller{
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
	}
	todoRoutes := router.Group("/todo")
	todoRoutes.Get("/:id", h.get)
	todoRoutes.Delete("/:id", h.delete)
	todoRoutes.Patch("/:id", h.update)
	todoRoutes.Post("/:id", h.add)
}

func (t controller) add(c *fiber.Ctx) error {
	var todo Todo
	if err := rest.ParseBodyAndValidate(c, &todo); err != nil {
		return c.JSON(err)
	}
	if err := t.todoService.db.Create(&todo); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Unable to add todo",
		})
	}
	return c.JSON(fiber.Map{
		"status": 1,
	})
}

func (t controller) update(c *fiber.Ctx) error {
	return c.JSON("WIP")

}

func (t controller) get(c *fiber.Ctx) error {
	return c.JSON("WIP")

}

func (t controller) delete(c *fiber.Ctx) error {
	return c.JSON("WIP")
}
