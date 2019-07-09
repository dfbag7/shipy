package main

import (
	"fmt"
	pb "github.com/dfbag7/shipy/user-service/proto/auth"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"log"
)

func main() {
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}
	
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shipy.auth"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Will comment this out for now to save having to run this locally...
	publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
