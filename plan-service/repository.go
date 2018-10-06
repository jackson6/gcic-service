// plan-service/repository.go

package main

import (
	pb "github.com/jackson6/gcic-service/plan-service/proto/plan"
	"gopkg.in/mgo.v2"
)

const (
	dbName = "invest"
	planCollection = "plan"
)

type Repository interface {
	Create(*pb.Plan) (*pb.Plan, error)
	All(*pb.Request) ([]*pb.Plan, error)
	Get(*pb.Plan) (*pb.Plan, error)
	Delete(*pb.Plan) error
	Update(*pb.Plan) error
	Close()
}

type PlanRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *PlanRepository) Create(plan *pb.Plan)  (*pb.Plan, error) {
	err := repo.collection().Insert(plan)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

// GetAll users
func (repo *PlanRepository) All(request *pb.Request) ([]*pb.Plan, error) {
	var plans []*pb.Plan
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&plans)
	return plans, err
}

func (repo *PlanRepository) Get(plan *pb.Plan) (*pb.Plan, error) {
	err := repo.collection().FindId(plan.Id).One(&plan)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (repo *PlanRepository) Delete(plan *pb.Plan) (error) {
	err := repo.collection().Remove(plan.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PlanRepository) Update(plan *pb.Plan) (error) {
	err := repo.collection().Update(plan.Id, &plan)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database session after each query has ran.
// Mgo creates a 'master' session on start-up, it's then good practice
// to copy a new session for each request that's made. This means that
// each request has its own database session. This is safer and more efficient,
// as under the hood each session has its own database socket and error handling.
// Using one main database socket means requests having to wait for that session.
// I.e this approach avoids locking and allows for requests to be processed concurrently. Nice!
// But... it does mean we need to ensure each session is closed on completion. Otherwise
// you'll likely build up loads of dud connections and hit a connection limit. Not nice!
func (repo *PlanRepository) Close() {
	repo.session.Close()
}

func (repo *PlanRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(planCollection)
}
