package api

import (
	db "github.com/SaifAlqady51/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer function creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Accounts Endpoints
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// Transfers endpoints
	router.POST("/transfers", server.createTransfer)

	// Users endpoints
	router.POST("/users", server.CreateUser)

	server.router = router

	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
