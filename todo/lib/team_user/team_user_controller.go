package team_user

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
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
	// teamRoutes.Get("/", h.getTeams)
	teamRoutes.Post(fmt.Sprintf("/:id/%s", h.config.ACCOUNT.PREFIX), h.addMembers)
	// teamRoutes.Get("/:id", h.getTeam)
	// teamRoutes.Patch("/:id", h.updateTeam)
	// teamRoutes.Delete("/:id", h.deleteTeam)
}

func (h controller) addMembers(c *fiber.Ctx) error {
	teamID, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	users := strings.Split(string(c.Request().Body()), ",")

	teamUsers := *new([]TeamUser)
	for _, u := range users {
		userID, err := strconv.ParseUint(u, 0, 0)
		if err != nil {
			return c.JSON(fiber.ErrBadRequest)
		}
		t := TeamUser{
			TeamID: uint(teamID),
			UserID: uint(userID),
		}
		teamUsers = append(teamUsers, t)
	}

	if err := h.teamService.db.Model(&TeamUser{}).Create(&teamUsers).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrBadRequest)
	}

	return c.JSON(teamUsers)
}
