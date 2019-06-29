package main

import (
	pb "github.com/dfbag7/shipy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "shipy"
	vesselCollection = "vessels"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselMongoRepository struct {
	session *mgo.Session
}

func (repo *VesselMongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define a more complex query than our consignment-serve's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.
	err := repo.collection().Find(bson.M {
		"capacity": bson.M{ "$gte": spec.Capacity },
		"maxweight": bson.M{ "$gte": spec.MaxWeight },
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselMongoRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselMongoRepository) Close() {
	repo.session.Close()
}

func (repo *VesselMongoRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
