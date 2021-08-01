package team_user

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

func (r *Repository) GetTeamsOfUser(teamUsers []TeamUser, userId uint) *gorm.DB {
	return r.db.Where("user_id=?", userId).Find(&teamUsers)
}
