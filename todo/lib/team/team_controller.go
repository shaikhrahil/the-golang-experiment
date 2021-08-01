package team

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/user"
	"gorm.io/gorm"
)

type controller struct {
	logger      *log.Logger
	teamService Repository
	config      rest.Configuration
}

func NewController(r *fiber.Router, conf rest.Configuration, logger *log.Logger, accountService Repository) {
	router := *r
	h := controller{
		logger:      logger,
		teamService: accountService,
		config:      conf,
	}
	teamRoutes := router.Group(h.config.TEAM.PREFIX)
	teamRoutes.Get("/", h.getTeams)
	teamRoutes.Post("/", h.addTeam)
	teamRoutes.Get("/:id", h.getTeam)
	teamRoutes.Patch("/:id", h.updateTeam)
	teamRoutes.Post("/:id/users", h.addMembers)
	teamRoutes.Delete("/:id", h.deleteTeam)
}

func (u *controller) addTeam(c *fiber.Ctx) error {
	var team Team

	if err := rest.ParseBodyAndValidate(c, &team); err != nil {
		u.logger.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	userID := rest.GetUser(c)
	team.Users = append(team.Users, user.User{
		User: accounts.User{
			Model: rest.Model{
				ID: userID,
			},
			FirstName: "slkmslmka",
		},
	})

	res := u.teamService.db.Create(&team)

	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: res.Error.Error(),
		})
	}
	return c.JSON(team)
}

func (h controller) getTeam(c *fiber.Ctx) error {
	teamId := c.Params("id")
	var team Team
	res := h.teamService.db.Where("id = ?", teamId).First(&team)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.JSON(team)
}

func (h controller) getTeams(c *fiber.Ctx) error {
	var teams []Team
	res := h.teamService.db.Find(&teams)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})

	}
	return c.JSON(teams)
}

func (h controller) deleteTeam(c *fiber.Ctx) error {
	teamId := c.Params("id")
	if res := h.teamService.db.Delete(&Team{}, teamId); res.Error != nil {
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
	if res := h.teamService.db.Where("id = ?", teamId).Updates(team); res.Error != nil {
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

func (h controller) addMembers(c *fiber.Ctx) error {
	// teamId := c.Params("id")
	users := strings.Split(string(c.Request().Body()), ",")
	var updatedTeam Team
	var dbUsers []user.User

	if err := h.teamService.db.Model(&updatedTeam).Association("Users").Find(&dbUsers, users); err != nil {
		h.logger.Println(err.Error())
		return c.JSON(err.Error())
	}

	return c.JSON(updatedTeam)
}
