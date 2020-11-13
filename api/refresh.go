package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Oloruntobi1/bankBackend/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (server *Server) refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "name me dey sup 1")
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	// os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	gotToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		fmt.Println("the error: ", err)
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := gotToken.Claims.(jwt.Claims); !ok && !gotToken.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := gotToken.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && gotToken.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, "na me dey sup 2")
			return
		}

		userEmail := fmt.Sprintf("%v", claims["user_id"])
		// userId, err := strconv.Parse(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		// if err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, "Error occurred")
		// 	return
		// }
		//Delete the previous Refresh Token
		deleted, delErr := token.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		// get the client
		client, err := server.store.GetClient(c, userEmail)

		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		//Create new pairs of refresh and access tokens
		ts := token.PrepareToken(client)

		//save the tokens metadata to redis
		saveErr := token.CreateAuth(client, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
