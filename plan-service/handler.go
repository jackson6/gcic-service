// plan-service/handler.go

package main

import (
	pb "github.com/jackson6/gcic-service/plan-service/proto/plan"
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
	return &PlanRepository{s.session.Clone()}
}

// CreateUser - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) Create(ctx context.Context, req *pb.Plan, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	plan, err := repo.Create(req)
	if err != nil {
		return err
	}
	res.Plan = plan
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.Plan, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	plan, err := repo.Get(req)
	if err != nil {
		return err
	}
	res.Plan = plan
	return nil
}

func (s *service) All(ctx context.Context,req *pb.Request, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	plans, err := repo.All(req)
	if err !=nil {
		return err
	}
	res.Plans = plans
	return nil
}

func (s *service) Delete(ctx context.Context,req *pb.Plan, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Delete(req)
	if err !=nil {
		return err
	}
	res.Code = 200
	return nil
}

func (s *service) Update(ctx context.Context,req *pb.Plan, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()


	err := repo.Update(req)
	if err !=nil {
		return err
	}
	res.Code = 200
	return nil
}