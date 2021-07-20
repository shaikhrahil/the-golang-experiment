package auth

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
	"github.com/shaikhrahil/the-golang-experiment/rest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type controller struct {
	logger         *log.Logger
	authService    Repository
	accountService accounts.Repository
}

func NewController(r *fiber.Router, logger *log.Logger, authService Repository, accountsService accounts.Repository) {
	router := *r
	h := controller{
		logger:         logger,
		authService:    authService,
		accountService: accountsService,
	}
	authRoutes := router.Group("/auth")
	authRoutes.Post("/login", h.Login)
	authRoutes.Post("/signup", h.Signup)
	authRoutes.Post("/logout", h.Logout)
	authRoutes.Post("/token/refresh", h.RefreshToken)
}

type loginReq struct {
	email    string `validate:"required,email,min=3,max=32" partial_validate:"omitempty,email,min=3,max=32"`
	password string `validate:"required,min=10,max=32"`
}

func (u controller) Login(c *fiber.Ctx) error {
	var user loginReq
	var userDB accounts.User
	if err := rest.ParseAndValidatePartially(c, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	res := u.accountService.GetByEmail(&userDB, user.email)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: `Invalid Username or EmailID`,
		})
	}
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.password), 10)
	if hashErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: `Invalid Password`,
		})
	}
	user.password = string(hash)

	if err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(userDB.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: `Invalid Password`,
		})

	}

	tkn := Generate(&TokenPayload{
		ID: userDB.ID,
	})

	return c.JSON(fiber.Map{
		"token": tkn,
	})
}

func (u *controller) Signup(c *fiber.Ctx) error {
	var user accounts.User
	if err := rest.ParseBodyAndValidate(c, user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)

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
	})

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
