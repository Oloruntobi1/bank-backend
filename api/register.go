package api

import (
	"net/http"

	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/Oloruntobi1/bankBackend/helper"
	"github.com/Oloruntobi1/bankBackend/token"
	"github.com/gin-gonic/gin"
)



type createRegistrationRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type createRegistrationResponse struct {
	Client db.Client
	Token string
}

func(server *Server) register(ctx *gin.Context) {

	var req createRegistrationRequest
	var res createRegistrationResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	password := helper.HashAndSalt([]byte(req.Password))
	arg := db.CreateClientParams {
		Name: req.Name,
		Email: req.Email,
		Password: password,
	}

	client, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res.Client = client
	res.Token = token.PrepareToken(client)
	
	ctx.JSON(http.StatusOK, res)

}