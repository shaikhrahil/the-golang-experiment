package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
)

type controller struct {
	logger         *log.Logger
	authService    Repository
	accountService accounts.Repository
	config         rest.Configuration
}

func NewController(r *fiber.Router, conf rest.Configuration, logger *log.Logger, authService Repository, accountsService accounts.Repository) {
	router := *r
	h := controller{
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
		config:         conf,
	}
	authRoutes := router.Group(h.config.AUTH.PREFIX)
	authRoutes.Post("/login", h.Login)
	authRoutes.Post("/signup", h.Signup)
	authRoutes.Post("/logout", h.Logout)
	authRoutes.Post("/token/refresh", h.RefreshToken)
}

type loginReq struct {
	Email    string `validate:"required,email,min=3,max=32" partial_validate:"omitempty,email,min=3,max=32"`
	Password string `validate:"required,min=10,max=32"`
}

func (u controller) Login(c *fiber.Ctx) error {
	var creds loginReq
	var userDB accounts.User
	if err := rest.ParseBodyAndValidate(c, &creds); err != nil {
		u.logger.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	res := u.accountService.GetByEmail(&userDB, creds.Email)

	if res.Error != nil {
		u.logger.Println(res.Error.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: `Invalid Username or EmailID`,
		})
	}

	if !CheckPasswordHash(creds.Password, userDB.Password) {
		u.logger.Println("Invallid password")
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: `Invalid Password`,
		})
	}

	tkn := Generate(&TokenPayload{
		ID: userDB.ID,
	}, u.config.AUTH.JWT_TTL, u.config.AUTH.JWT_SECRET)

	return c.JSON(fiber.Map{
		"token": tkn,
	})
}

func (u *controller) Signup(c *fiber.Ctx) error {
	var user accounts.User

	if err := rest.ParseBodyAndValidate(c, &user); err != nil {
		u.logger.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	res := u.accountService.AddUser(&user)

	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: res.Error.Error(),
		})
	}

	tkn := Generate(&TokenPayload{
		ID: user.ID,
	}, u.config.AUTH.JWT_TTL, u.config.AUTH.JWT_SECRET)

	return c.JSON(fiber.Map{
		"token": tkn,
	})

}

func (u *controller) Logout(c *fiber.Ctx) error {
	return c.JSON("Workinh !!")
}

func (u *controller) RefreshToken(c *fiber.Ctx) error {
	return c.JSON("Workinh !!")
}
