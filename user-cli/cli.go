package main

import (
	pb "github.com/dfbag7/shipy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main() {
	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Define our flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag {
				Name: "name",
				Usage: "User's full name",
			},
			cli.StringFlag {
				Name: "email",
				Usage: "User's email",
			},
			cli.StringFlag {
				Name: "password",
				Usage: "User's password",
			},
			cli.StringFlag {
				Name: "company",
				Usage: "User's company name",
			},
		))

	// Start as service
	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User {
				Name: name,
				Email: email,
				Password: password,
				Company: company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
