package team

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type controller struct {
	logger         *log.Logger
	accountService Repository
	config         rest.Configuration
}

func NewController(r *fiber.Router, conf rest.Configuration, logger *log.Logger, accountService Repository) {
	router := *r
	h := controller{
		logger:         logger,
		accountService: accountService,
		config:         conf,
	}
	accountRoutes := router.Group(h.config.ACCOUNT.PREFIX)
	accountRoutes.Get("/", h.getTeams)
	accountRoutes.Get("/:id", h.getTeam)
	accountRoutes.Patch("/:id", h.updateTeam)
	accountRoutes.Delete("/:id", h.deleteTeam)
}

func (h controller) getTeam(c *fiber.Ctx) error {
	teamId := c.Params("id")
	var team Team
	res := h.accountService.db.Where("id = ?", teamId).First(&team)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.JSON(team)
}

func (h controller) getTeams(c *fiber.Ctx) error {
	var teams []Team
	res := h.accountService.db.Find(&teams)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	}
	return c.JSON(teams)
}

func (h controller) deleteTeam(c *fiber.Ctx) error {
	teamId := c.Params("id")
	if res := h.accountService.db.Delete(&Team{}, teamId); res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	} else if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})

	}
	return c.JSON(teamId)

}

func (h controller) updateTeam(c *fiber.Ctx) error {
	var team Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := rest.ValidateStructPartially(team)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	teamId := c.Params("id")
	if res := h.accountService.db.Where("id = ?", teamId).Updates(team); res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": gorm.ErrRecordNotFound.Error(),
		})
	}

	return c.JSON(teamId)
}
