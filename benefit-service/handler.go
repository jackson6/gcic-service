// benefit-service/handler.go

package main

import (
	"encoding/json"
	pb "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &BenefitRepository{s.session.Clone()}
}

func converter(data interface{}, dataType int) (interface{}, error){
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if dataType == 0 {
		benefit := new(pb.Benefit)
		err = json.Unmarshal(jsonData, &benefit)
		if err != nil {
			return nil, err
		}
		return benefit, nil
	} else {
		benefits := make([]*pb.Benefit, 0)
		err = json.Unmarshal(jsonData, &benefits)
		if err != nil {
			return nil, err
		}
		return benefits, nil
	}
}

// CreateUser - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) Create(ctx context.Context, req *pb.Benefit, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	benefit, err := repo.Create(req)
	if err != nil {
		return err
	}
	res.Benefit = benefit
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.Benefit, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.Get(req)
	if err != nil {
		return err
	}

	benefit, err := converter(data, 0)
	if err != nil {
		return err
	}
	res.Benefit = benefit.(*pb.Benefit)
	return nil
}

func (s *service) All(ctx context.Context,req *pb.Request, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.All(req)
	if err !=nil {
		return err
	}

	benefits, err := converter(data, 1)
	if err != nil {
		return err
	}

	res.Benefits = benefits.([]*pb.Benefit)
	return nil
}

func (s *service) Delete(ctx context.Context,req *pb.Benefit, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Delete(req)
	if err !=nil {
		return err
	}
	res.Code = 200
	return nil
}

func (s *service) Update(ctx context.Context,req *pb.Benefit, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()


	err := repo.Update(req)
	if err !=nil {
		return err
	}
	res.Code = 200
	return nil
}