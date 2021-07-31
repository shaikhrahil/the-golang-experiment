package auth

import (
	"log"

	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
)

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
	config Config
}

type Config struct {
	Claims func(payload accounts.User) jwt.MapClaims
}

func NewRepository(db *gorm.DB, logger *log.Logger) Repository {
	return Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) Generate(payload accounts.User, expiresIn string, secret string) string {
	p := jwt.MapClaims{
		"UserID": payload.ID,
	}
	if r.config.Claims != nil {
		p = r.config.Claims(payload)
	}
	return generate(p, expiresIn, secret)
}

func (r Repository) SetConfig(config Config) Repository {
	r.config = config
	return r
}
