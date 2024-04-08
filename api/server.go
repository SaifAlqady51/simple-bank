package api

import (
	"fmt"

	db "github.com/SaifAlqady51/simple-bank/db/sqlc"
	"github.com/SaifAlqady51/simple-bank/token"
	"github.com/SaifAlqady51/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	tokenMaker token.Maker
	store      db.Store
	router     *gin.Engine
	config     util.Config
}

// NewServer function creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: current token length is %d %w", len(config.TokenSymmetricKey), err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// Accounts Endpoints
	authRoute.POST("/accounts", server.createAccount)
	authRoute.GET("/accounts/:id", server.getAccount)
	authRoute.GET("/accounts", server.listAccounts)

	// Transfers endpoints
	authRoute.POST("/transfers", server.createTransfer)

	// Users endpoints
	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.loginUser)

	server.router = router

	return server, nil

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
