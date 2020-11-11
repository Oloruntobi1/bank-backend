package token

import (
	"time"

	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/Oloruntobi1/bankBackend/helper"
	"github.com/dgrijalva/jwt-go"
)


func PrepareToken(client db.Client) string {
	tokenContent := jwt.MapClaims{
		"user_id": client.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helper.HandleErr(err)

	return token
}