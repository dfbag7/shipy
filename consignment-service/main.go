// shipy-service-consignment/main.go

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/dfbag7/shipy/consignment-service/proto/consignment"
	userService "github.com/dfbag7/shipy/user-service/proto/user"
	vesselProto "github.com/dfbag7/shipy/vessel-service/proto/vessel"
)

const (
	defaultHost = "mongodb://datastore:27017"
)

func main() {
	// Set-up micro service
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("consignment"),
		micro.Version("latest"),

		// Our auth middleware
		micro.WrapHandler(AuthWrapper),
	)

	// Init will parse the command line flags.
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shipy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shipy.vessel.service", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	log.Print("Run the consignment service")

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure whi this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})

		log.Println("Auth resp: ", authResp)
		log.Println("Err: ", err)

		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)

		return err
	}
}
