package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	auth "github.com/shaikhrahil/the-golang-experiment/auth/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team_user"
	"gorm.io/gorm"
)

type controller struct {
	config         rest.Configuration
	logger         *log.Logger
	accountService accounts.Repository
	authService    auth.Repository
	todoService    Repository
}

func NewController(r *fiber.Router, conf rest.Configuration, logger *log.Logger, authService auth.Repository, accountsService accounts.Repository, todoService Repository) {
	router := *r
	h := controller{
		config:         conf,
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
		todoService:    todoService,
	}
	todoRoutes := router.Group(h.config.TODO.PREFIX)
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

	var teamUser rest.MapModel
	userID := rest.GetUser(c)
	teamID := rest.GetTeams(c)
	if err := h.todoService.db.Model(&team_user.TeamUser{}).Where("team_id = ? and user_id = ?", teamID, userID).First(&teamUser).Error; err != nil {
		h.logger.Println(err.Error())
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Not found in team_users",
		})
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

func (h controller) getAll(c *fiber.Ctx) error {
	var todos []TodoSummary
	if err := h.todoService.db.Model(&Todo{}).Find(&todos, &Todo{TeamUserID: rest.GetUser(c)}).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrNotFound)
	}
	return c.JSON(todos)
}

func (h controller) get(c *fiber.Ctx) error {
	var todo Todo
	todoID := c.Params("id")
	if err := h.todoService.db.First(&todo, todoID).Error; err != nil {
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
