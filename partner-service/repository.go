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
	All() ([]*Partner, error)
	Get(*pb.Partner) (*Partner, error)
	Delete(*pb.Partner) error
	Update(*pb.Partner) error
	Close()
}

type PartnerRepository struct {
	session *mgo.Session
}

type Partner struct {
	Id  bson.ObjectId `bson:"_id"json:"id"`
	Name string `bson:"name"json:"name"`
	Address string `bson:"address"json:"address"`
	Parish string `bson:"parish"json:"parish"`
	Country string `bson:"country"json:"country"`
	Contact string `bson:"contact"json:"contact"`
	Logo string `bson:"logo"json:"logo"`
	Coord Coord `bson:"coord"json:"coord"`
}

type Coord struct {
	Lat float64 `bson:"lat"json:"lat"`
	Lon float64 `bson:"long"json:"long"`
}

// Create a new user
func (repo *PartnerRepository) Create(partner *pb.Partner)  (*pb.Partner, error) {
	err := repo.collection().Insert(partner)
	if err != nil {
		return nil, err
	}
	return partner, nil
}

// GetAll users
func (repo *PartnerRepository) All() ([]*Partner, error) {
	var partners []*Partner
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&partners)
	return partners, err
}

func (repo *PartnerRepository) Get(partner *pb.Partner) (*Partner, error) {
	newPartner := new(Partner)
	err := repo.collection().FindId(bson.ObjectIdHex(partner.Id)).One(&newPartner)
	if err != nil {
		return nil, err
	}
	return newPartner, nil
}

func (repo *PartnerRepository) Delete(partner *pb.Partner) (error) {
	err := repo.collection().RemoveId(bson.ObjectIdHex(partner.Id))
	if err != nil {
		return err
	}
	return nil
}

func (repo *PartnerRepository) Update(partner *pb.Partner) (error) {
	err := repo.collection().Update(bson.ObjectIdHex(partner.Id), &partner)
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
