package main

import (
	"context"
	pb "github.com/dfbag7/shipy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselMongoRepository struct {
	collection *mongo.Collection
}

func (repo *VesselMongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D {{
		"capacity",
		bson.D{ {
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}

	var vessel *pb.Vessel
	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}
