package todo

import (
	"github.com/dgrijalva/jwt-go"
	accounts "github.com/shaikhrahil/the-golang-experiment/accounts/lib"
)

func NewClaims(payload accounts.User) jwt.MapClaims {
	p := jwt.MapClaims{
		"UserID": payload.ID,
	}
	return p
}
