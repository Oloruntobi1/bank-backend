package api

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )


type loginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// func (server *Server) login(ctx *gin.Context) {

// 	var req loginRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

// 	// if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
// 	// 	return map[string]interface{}{"message": "Wrong password"}
// 	// }

// }