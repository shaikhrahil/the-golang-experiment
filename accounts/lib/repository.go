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

func (r *Repository) GetByEmail(user *User, email string) *gorm.DB {
	return r.db.Find(&user).Where("email = ?", email)
}

func (r *Repository) GetByUserName(user *User, userName string) *gorm.DB {
	return r.db.Find(&user).Where("userName = ?", userName)
}

func (r *Repository) AddUser(user *User) *gorm.DB {
	return r.db.Create(&user)
}
