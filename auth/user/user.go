package user

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
}

func AddHandlers(r fiber.Router) {
	handler := r.Group("/user")
	handler.Get("/:id?", GetUser)
	handler.Post("", GetUser)
	handler.Patch("", GetUser)
	handler.Put("", GetUser)
	handler.Delete("", GetUser)
}

func SignUp() {

}
func Login() {

}
func Logout() {

}
func DeleteAccount() {

}

func GetUser(c *fiber.Ctx) {
	user := new(User)
	if userId := c.Params("id"); userId != "" {
		user.Name = userId
	}
	c.JSON(user)
}
