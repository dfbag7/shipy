// shippy-cli-consignment/main.go
package main

import (
	"encoding/json"
	"errors"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/dfbag7/shipy/consignment-service/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}

	return consignment, err
}

func main() {

	err := cmd.Init()
	if err != nil {
		log.Fatalf("Could not init: %v", err)
	}

	// Create new client
	client := pb.NewShippingServiceClient("consignment", microclient.DefaultClient)

	// Contact the server and print our its response
	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecitng file and token."))
	}

	log.Println(os.Args)

	file := os.Args[1]
	token := os.Args[2]

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// First call using our tokenised context
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created consignment: %t", r.Created)

	// Second call
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}


