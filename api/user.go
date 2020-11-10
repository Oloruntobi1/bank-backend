package api

import (
	"net/http"
	db "github.com/Oloruntobi1/Oloruntobi1/bank_backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/Oloruntobi1/Oloruntobi1/bank_backend/helper"
)

type createUserRequest struct {
	Name   string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func(server *Server) createUser(ctx *gin.Context) {
	
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	password := helper.HashAndSalt([]byte(req.Password))
	arg := db.CreateUserParams {
		Name: req.Name,
		Email: req.Email,
		Password: password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}