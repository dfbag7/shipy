package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/dfbag7/shipy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "datastore:27017"
)

func createDummyData(repo repository) {
	vessels := []*pb.Vessel {
		{ Id: "vessel1001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500 },
	}
	for _,v := range vessels {
		repo.Create(v)
	}
}


func main() {
	mongoHost := os.Getenv("DB_HOST")
	if mongoHost == "" {
		mongoHost = defaultHost
	}

	session, err := CreateSession(mongoHost)
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}

	defer session.Close()

	repo := &VesselMongoRepository{ session.Copy() }

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("shipy.vessel.service"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
