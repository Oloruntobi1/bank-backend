package api

import (
	"github.com/mcuadros/go-gin-prometheus"

	db "github.com/Oloruntobi1/bankBackend/db/sqlc"
	"github.com/Oloruntobi1/bankBackend/token"
	"github.com/gin-gonic/gin"
)

// Server struct to hold our store and the gin router
type Server struct {
	store *db.Store
	router *gin.Engine
}

//NewServer contains our server routes 
func NewServer(store *db.Store) *Server {
	server := &Server{store : store}
	r := gin.New()

	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)

	// r.Run(":29090")

	
	r.POST("/users", token.TokenAuthMiddleware(), server.createUser)
	r.DELETE("/user/:id", token.TokenAuthMiddleware(), server.deleteUser)
	r.POST("/register", server.register)
	r.POST("/login", token.TokenAuthMiddleware(), server.login)
	r.POST("/logout", token.TokenAuthMiddleware(), server.logout)
	r.POST("/refresh", token.TokenAuthMiddleware(), server.refresh)
	
	server.router = r
	return server
}

// Start helps start a new server
func (server *Server) Start (address string) error {

	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}