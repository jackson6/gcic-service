package main

import (
	pb "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"github.com/stripe/stripe-go"
	stripeCharge "github.com/stripe/stripe-go/charge"
	stripeCustomer "github.com/stripe/stripe-go/customer"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2"
	"os"
)

const (
	dbName = "invest"
	userCollection = "transaction"
)

type Repository interface {
	CreateCharge(charge *pb.Charge) (*stripe.Charge, error)
	CreateCustomer(charge *pb.Customer) (*stripe.Customer, error)
	Close()
}

type PaymentRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *PaymentRepository) CreateTransaction(transaction *pb.Transaction) error {
	return repo.collection().Insert(transaction)
}

func (repo *PaymentRepository) CreateCharge(charge *pb.Charge) (*stripe.Charge, error) {
	stripe.Key = os.Getenv("STRIPE_KEY")

	args := &stripe.ChargeParams {
		Amount: stripe.Int64(charge.Amount),
		Currency: stripe.String(string(charge.Currency)),
		Description: stripe.String(charge.Description),
	}

	if charge.Customer != "" {
		args.Customer = stripe.String(charge.Customer)
	} else {
		args.SetSource(charge.Token) // obtained with Stripe.js
	}

	u2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	args.SetIdempotencyKey(u2.String())

	ch, err := stripeCharge.New(args)
	if err != nil {
		return nil, err
	}
	transaction := &pb.Transaction{
		UserId: charge.UserId,
		ChargeId: ch.ID,
		Amount: ch.Amount,
		Currency: string(ch.Currency),
		Description: ch.Description,
		IdempotencyKey: u2.String(),
	}
	err = repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (repo *PaymentRepository) CreateCustomer(customer *pb.Customer) (*stripe.Customer, error) {
	stripe.Key = os.Getenv("STRIPE_KEY")
	customerParams := &stripe.CustomerParams{
		Email: stripe.String(customer.Email),
	}
	customerParams.SetSource(customer.Token)

	newCustomer, err := stripeCustomer.New(customerParams)
	if err != nil {
		return nil, err
	}
	return newCustomer, nil
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
func (repo *PaymentRepository) Close() {
	repo.session.Close()
}

func (repo *PaymentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(userCollection)
}