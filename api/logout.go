package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Oloruntobi1/bankBackend/token"

)

func (server *Server) logout(c *gin.Context) {
	metadata, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := token.DeleteTokens(metadata)
	if delErr != nil {
		c.JSON(http.StatusUnauthorized, delErr.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}