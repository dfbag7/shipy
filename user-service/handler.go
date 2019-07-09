package main

import (
	"encoding/json"
	"errors"
	pb "github.com/dfbag7/shipy/user-service/proto/user"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

const topic = "user.created"

type service struct {
	repo Repository
	tokenService Authable
	PubSub broker.Broker
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with: ", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	log.Println(user)

	// Compare our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token

	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Generate a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}

	res.User = req
	if err := srv.publishEvent(req); err != nil {
		return err
	}

	return nil
}

func (srv *service) publishEvent(user *pb.User) error {
	// Marshal to JSON string
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Create a broker message
	msg := &broker.Message{
		Header: map[string]string {
			"id": user.Id,
		},
		Body: body,
	}

	// Publish message to broker
	// and do not return error if the publish operation is failed
	if err := srv.PubSub.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	} else {
		log.Println("Published event")
	}

	return nil
}


func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	log.Printf("Validating token: %s", req.Token)

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
