package team

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

// func (r *Repository) GetByEmail(team *Team, email string) *gorm.DB {
// 	return r.db.Where("email = ?", email).First(&team)
// }

// func (r *Repository) GetByTeamName(team *Team, teamName string) *gorm.DB {
// 	return r.db.Find(&team).Where("teamName = ?", teamName)
// }

// func (r *Repository) Addteam(team *Team) *gorm.DB {
// 	return r.db.Create(&team)
// }
