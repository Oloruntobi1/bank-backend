package token

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/Oloruntobi1/bankBackend/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"github.com/Oloruntobi1/bankBackend/rdstore"
)


type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func PrepareToken(client db.Client) *TokenDetails {

	var err error 
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(client.ID))
	
	atContent := jwt.MapClaims{
		"user_id": client.ID,
		"expiry": td.AtExpires,
		"access_uuid": td.AccessUuid,
		"authorized": true,

	}
	aToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), atContent)
	td.AccessToken, err = aToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	helper.HandleErr(err)

		
	rtContent := jwt.MapClaims{
		"user_id": client.ID,
		"expiry": td.RtExpires,
		"access_uuid": td.RefreshUuid,
		"authorized": true,

	}
	rToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), rtContent)
	td.RefreshToken, err = rToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	helper.HandleErr(err)

	return td
}


func CreateAuth(client db.Client, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := rdstore.Client.Set(td.AccessUuid, strconv.Itoa(int(client.ID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := rdstore.Client.Set(td.RefreshUuid, strconv.Itoa(int(client.ID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   err := TokenValid(c.Request)
	   if err != nil {
		  c.JSON(http.StatusUnauthorized, err.Error())
		  c.Abort()
		  return
	   }
	   c.Next()
	}
  }

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}