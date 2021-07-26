package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type controller struct {
	logger         *log.Logger
	accountService accounts.Repository
	authService    auth.Repository
	todoService    Repository
}

func NewController(r *fiber.Router, logger *log.Logger, authService auth.Repository, accountsService accounts.Repository, todoService Repository) {
	router := *r
	h := controller{
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
		todoService:    todoService,
	}
	todoRoutes := router.Group("/todo")
	todoRoutes.Get("/:id", h.get)
	todoRoutes.Delete("/:id", h.delete)
	todoRoutes.Patch("/:id", h.update)
	todoRoutes.Post("/", h.add)
}

func (h controller) add(c *fiber.Ctx) error {
	var todo Todo
	if err := rest.ParseBodyAndValidate(c, &todo); err != nil {
		return c.JSON(err)
	}
	if err := h.todoService.db.Create(&todo).Error; err != nil {
		h.logger.Println(err.Error())
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Unable to add todo",
		})
	}
	return c.JSON(todo)
}

func (h controller) update(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := rest.ValidateStructPartially(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	todoId := c.Params("id")
	if res := h.todoService.db.Where("id = ?", todoId).Updates(todo); res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})
	}

	return c.JSON(todo)
}

func (h controller) get(c *fiber.Ctx) error {
	var todos []Todo
	var user User
	// todoID := c.Params("id")
	if err := h.todoService.db.Model(&user).Association("Todos").Find(&todos); err != nil {
		// if err := h.todoService.db.Preload("todo", "todo_id = ?", todoID).First(&todo).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(todos)
}

func (h controller) delete(c *fiber.Ctx) error {
	var todo Todo
	todoID := c.Params("id")
	if err := h.todoService.db.Where("id = ?", todoID).Delete(&todo).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(todo)
}
