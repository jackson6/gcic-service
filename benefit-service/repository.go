// benefit-service/repository.go

package main

import (
	pb "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	benefitCollection = "benefits"
)

type Repository interface {
	Create(*pb.Benefit) (*pb.Benefit, error)
	All(*pb.Request) ([]*Benefit, error)
	Get(*pb.Benefit) (*Benefit, error)
	Delete(*pb.Benefit) error
	Update(*pb.Benefit) error
	Close()
}

type Benefit struct {
	Id bson.ObjectId `bson:"_id"json:"id"`
	Title string `bson:"title"json:"title"`
	Description string `bson:"description"json:"description"`
	PartnerId string `bson:"partnerid"json:"partner_id"`
	Img []string `bson:"img"json:"img"`
}

type BenefitRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *BenefitRepository) Create(benefit *pb.Benefit)  (*pb.Benefit, error) {
	err := repo.collection().Insert(benefit)
	if err != nil {
		return nil, err
	}
	return benefit, nil
}

// GetAll users
func (repo *BenefitRepository) All(request *pb.Request) ([]*Benefit, error) {
	var benefits []*Benefit
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&benefits)
	return benefits, err
}

func (repo *BenefitRepository) Get(benefit *pb.Benefit) (*Benefit, error) {
	data := new(Benefit)
	err := repo.collection().FindId(bson.ObjectIdHex(benefit.Id)).One(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *BenefitRepository) Delete(benefit *pb.Benefit) (error) {
	err := repo.collection().RemoveId(bson.ObjectIdHex(benefit.Id))
	if err != nil {
		return err
	}
	return nil
}

func (repo *BenefitRepository) Update(benefit *pb.Benefit) (error) {
	err := repo.collection().UpdateId(bson.ObjectIdHex(benefit.Id), &benefit)
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
func (repo *BenefitRepository) Close() {
	repo.session.Close()
}

func (repo *BenefitRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(benefitCollection)
}
