// partner-service/repository.go

package main

import (
	pb "github.com/jackson6/gcic-service/partner-service/proto/partner"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	partnerCollection = "partners"
)

type Repository interface {
	Create(*pb.Partner) (*pb.Partner, error)
	All() ([]*pb.Partner, error)
	Get(*pb.Partner) (*pb.Partner, error)
	Delete(*pb.Partner) error
	Update(*pb.Partner) error
	Close()
}

type PartnerRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *PartnerRepository) Create(partner *pb.Partner)  (*pb.Partner, error) {
	id := bson.NewObjectId()
	partner.Id = &id
	err := repo.collection().Insert(partner)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

// GetAll users
func (repo *PartnerRepository) All() ([]*pb.Partner, error) {
	var partners []*pb.Partner
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&partners)
	return partners, err
}

func (repo *PartnerRepository) Get(partner *pb.Partner) (*pb.Partner, error) {
	err := repo.collection().FindId(partner.Id).One(&partner)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

func (repo *PartnerRepository) Delete(partner *pb.Partner) (error) {
	err := repo.collection().Remove(partner.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PartnerRepository) Update(partner *pb.Partner) (error) {
	err := repo.collection().Update(partner.Id, &partner)
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
func (repo *PartnerRepository) Close() {
	repo.session.Close()
}

func (repo *PartnerRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(partnerCollection)
}
