package gapi

import (
	"fmt"

	db "github.com/SaifAlqady51/simple-bank/db/sqlc"
	"github.com/SaifAlqady51/simple-bank/pb"
	"github.com/SaifAlqady51/simple-bank/token"
	"github.com/SaifAlqady51/simple-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	tokenMaker token.Maker
	store      db.Store
	config     util.Config
}

// NewServer function creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: current token length is %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil

}
