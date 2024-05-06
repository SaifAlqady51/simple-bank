package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/SaifAlqady51/simple-bank/api"
	db "github.com/SaifAlqady51/simple-bank/db/sqlc"
	"github.com/SaifAlqady51/simple-bank/gapi"
	"github.com/SaifAlqady51/simple-bank/pb"
	"github.com/SaifAlqady51/simple-bank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// get env vairables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config ", err)
	}
	// connect to the database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("we can't connect to the database")
	}
	// new Store
	store := db.NewStore(conn)

	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	//create new grpc server
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start a server ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot start a grpc server ", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start a grpc server: ", err)
	}

}

func runGinServer(config util.Config, store db.Store) {

	// start a new Server
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start a server ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start a server ", err)
	}
}
