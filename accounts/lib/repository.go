package accounts

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetByEmail(email string) (User, error) {
	var user User
	if res := r.db.Find(&user).Where("email = ?", email); res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (r *Repository) GetByUserName(userName string) (User, error) {
	var user User
	if res := r.db.Find(&user).Where("userName = ?", userName); res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (r *Repository) GetSomething() string {
	return "heheheheheehh"
}
