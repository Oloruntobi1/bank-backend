package api

import (
	"net/http"
	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/Oloruntobi1/bankBackend/helper"
	"database/sql"
	"github.com/Oloruntobi1/bankBackend/token"
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

	metadata, err := token.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	_, err = token.FetchAuth(metadata)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
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

type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func(server *Server) deleteUser(ctx *gin.Context) {

	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	gottenUser, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = server.store.DeleteUser(ctx, gottenUser.ID )
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Successfully deleted")


}