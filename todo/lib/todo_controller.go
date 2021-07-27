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
	todoRoutes.Get("/", h.getAll)
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
	todo.UserTodo.UserID = rest.GetUser(c)
	// todo.UserTodo = UserTodo{
	// 	UserID: *rest.GetUser(c),
	// 	TodoID: todo.ID,
	// }
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

func (h controller) getAll(c *fiber.Ctx) error {
	var todos []TodoSummary
	if err := h.todoService.db.Model(&Todo{}).Joins("left join user_todos on todos.user_todo_id = user_todos.id").Where("user_todos.user_id = ?", rest.GetUser(c)).Find(&todos).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(todos)
}

func (h controller) get(c *fiber.Ctx) error {
	var todo Todo
	todoID := c.Params("id")
	if err := h.todoService.db.Where("id = ?", todoID).First(&todo).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(todo)
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
