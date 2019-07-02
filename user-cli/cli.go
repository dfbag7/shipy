package main

import (
	pb "github.com/dfbag7/shipy/user-service/proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"log"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "Testing123"
	company := "BBC"

	log.Println("Attempting to create user: %s %s %s %s", name, email, password, company)

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	log.Printf("Created user with id: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v", email, err)
	}

	log.Printf("Your access token is: %s", authResponse.Token)
}
