package main

import (
	"context"

	pb "github.com/dfbag7/shipy/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := s.repository.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}
