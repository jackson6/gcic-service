// user-service/repository.go

package main

import (
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	userCollection = "users"
)

type Repository interface {
	Create(*pb.User) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Get(*pb.User) (*pb.User, error)
	GetByEmail(*bson.M) (*pb.User, error)
	Close()
}

type UserRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *UserRepository) Create(user *pb.User)  (*pb.User, error) {
	err := repo.collection().Insert(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll users
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&users)
	return users, err
}

func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	err := repo.collection().FindId(bson.ObjectId(user.UserId)).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(find *bson.M) (*pb.User, error) {
	var user *pb.User
	err := repo.collection().Find(find).One(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
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
func (repo *UserRepository) Close() {
	repo.session.Close()
}

func (repo *UserRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(userCollection)
}
