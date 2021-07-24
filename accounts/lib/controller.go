package accounts

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type controller struct {
	logger         *log.Logger
	accountService Repository
}

func NewController(r *fiber.Router, logger *log.Logger, accountService Repository) {
	router := *r
	h := controller{
		logger:         logger,
		accountService: accountService,
	}
	accountRoutes := router.Group("/accounts")
	accountRoutes.Get("/", h.getAccounts)
	accountRoutes.Get("/:id", h.getAccount)
	accountRoutes.Patch("/:id", h.updateAccount)
	accountRoutes.Delete("/:id", h.deleteAccount)
}

func (h controller) getAccount(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user User
	res := h.accountService.db.Where("id = ?", userId).First(&user)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	}
	return c.JSON(user)
}

func (h controller) getAccounts(c *fiber.Ctx) error {
	var users []User
	res := h.accountService.db.Find(&users)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	}
	return c.JSON(users)
}

func (h controller) deleteAccount(c *fiber.Ctx) error {
	userId := c.Params("id")
	if res := h.accountService.db.Delete(&User{}, userId); res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	} else if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})

	}
	return c.JSON(userId)

}

func (h controller) updateAccount(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := rest.ValidateStructPartially(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	userId := c.Params("id")
	if res := h.accountService.db.Where("id = ?", userId).Updates(user); res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})
	}

	return c.JSON(userId)
}
