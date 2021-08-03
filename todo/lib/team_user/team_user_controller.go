package team_user

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"github.com/shaikhrahil/the-golang-experiment/todo/lib/team"
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
	teamRoutes.Get(fmt.Sprintf("/:friendID/%s", h.config.ACCOUNT.PREFIX), h.getCommonTeams)
	teamRoutes.Post(fmt.Sprintf("/:teamID/%s", h.config.ACCOUNT.PREFIX), h.addMembers)
	teamRoutes.Delete(fmt.Sprintf("/:teamID/%s", h.config.ACCOUNT.PREFIX), h.removeUsers)
}

func (h controller) addMembers(c *fiber.Ctx) error {
	teamID, err := strconv.ParseUint(c.Params("teamID"), 0, 0)
	if err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrBadRequest)
	}
	users := strings.Split(string(c.Request().Body()), ",")

	teamUsers := *new([]TeamUser)
	for _, u := range users {
		userID, err := strconv.ParseUint(u, 0, 0)
		if err != nil {
			h.logger.Println(err.Error())
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

func (h controller) removeUsers(c *fiber.Ctx) error {
	teamID, err := strconv.ParseUint(c.Params("teamID"), 0, 0)
	if err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrBadRequest)
	}

	teams := rest.GetTeams(c)
	if !teams.Contains(teamID) {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrUnauthorized)
	}

	userBody := string(c.Request().Body())

	if userBody == "" {
		h.logger.Println("No userids provided")
		return c.JSON(fiber.ErrBadRequest)
	}

	users := strings.Split(string(userBody), ",")

	teamUsers := *new([]TeamUser)

	for _, u := range users {
		userID, err := strconv.ParseUint(u, 0, 0)
		if err != nil {
			h.logger.Println(err.Error())
			return c.JSON(fiber.ErrBadRequest)
		}
		t := TeamUser{
			TeamID: uint(teamID),
			UserID: uint(userID),
		}
		teamUsers = append(teamUsers, t)
	}

	if err := h.teamService.db.Model(&TeamUser{}).Delete(&teamUsers).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrBadRequest)
	}

	return c.JSON(teamUsers)
}

func (h controller) getCommonTeams(c *fiber.Ctx) error {
	friendID, err := strconv.ParseUint(c.Params("friendID"), 0, 0)
	if err != nil {
		h.logger.Println(err.Error())
		return c.JSON(fiber.ErrBadRequest)
	}
	userID := rest.GetUser(c)

	var commonTeams []team.Team

	if err := h.teamService.db.Joins("inner join (team_users T1 inner join team_users T2 on T2.user_id = ? and T2.team_id = ?) on T1.user_id = ? and T1.team_id = ?", friendID, userID).Find(&commonTeams).Error; err != nil {
		h.logger.Println(err.Error())
		return c.JSON(gorm.ErrRecordNotFound.Error())
	}

	return c.JSON(commonTeams)
}
