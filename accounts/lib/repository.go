package accounts

import (
	"log"

	"github.com/shaikhrahil/the-golang-experiment/rest"
	"gorm.io/gorm"
)

type repository struct {
	rest.Controller
}

func Repository(db *gorm.DB, logger *log.Logger) {

}

func (r *repository) GetByEmail(email string) (User, error) {
	var user User
	if res := r.DB.Find(&user).Where("email = ?", email); res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (r *repository) GetByUserName(userName string) (User, error) {
	var user User
	if res := r.DB.Find(&user).Where("userName = ?", userName); res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
