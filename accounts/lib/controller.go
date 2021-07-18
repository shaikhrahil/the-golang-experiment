package accounts

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type controller struct {
	logger *log.Logger
	repos  repositories
}

func NewController(r *fiber.Router, logger *log.Logger, repos map[string]Repository) {
	router := *r
	nmr := repositories{
		Accounts: repos["Accounts"],
	}
	h := controller{logger: logger, repos: nmr}
	accountRoutes := router.Group("/accounts")
	accountRoutes.Get("/", h.getAccounts)
	accountRoutes.Get("/:id", h.getAccount)
	accountRoutes.Post("/signup", h.signUp)
	accountRoutes.Patch("/:id", h.updateAccount)
	accountRoutes.Delete("/:id", h.deleteAccount)
}

func (h controller) signUp(c *fiber.Ctx) {
	var user User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return
	}

	errors := rest.ValidateStruct(user)
	if errors != nil {
		c.Status(fiber.StatusBadRequest).JSON(errors)
		return
	}

	u := h.repos.Accounts.db.Create(&user)

	if u.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": u.Error.Error(),
		})
		return
	}

	if err := c.JSON(&user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}

func (h controller) getAccount(c *fiber.Ctx) {
	userId := c.Params("id")
	var user User
	res := h.repos.Accounts.db.First(&user, "id = ?", userId)
	if res.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
		return
	}
	if err := c.JSON(user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return
	}
}

func (h controller) getAccounts(c *fiber.Ctx) {
	var users []User
	res := h.repos.Accounts.db.Find(&users)
	if res.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
		return
	}

	if err := c.JSON(users); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return
	}
}

func (h controller) deleteAccount(c *fiber.Ctx) {
	if res := h.repos.Accounts.db.Delete(&User{}, c.Params("id")); res.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
		return
	} else if res.RowsAffected == 0 {
		c.JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})
		return
	}
	c.JSON(fiber.Map{
		"status": 1,
	})
}

func (h controller) updateAccount(c *fiber.Ctx) {
	var user User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return
	}

	errors := rest.ValidateStructPartially(user)
	if errors != nil {
		c.Status(fiber.StatusBadRequest).JSON(errors)
		return
	}

	if res := h.repos.Accounts.db.Where("id = ?", c.Params("id")).Updates(user); res.Error != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(fiber.Map{
		"status": 1,
	})
}
