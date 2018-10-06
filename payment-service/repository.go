package main

import (
	pb "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"github.com/jinzhu/gorm"
	"github.com/nu7hatch/gouuid"
	"github.com/stripe/stripe-go"
	stripeCharge "github.com/stripe/stripe-go/charge"
	stripeCustomer "github.com/stripe/stripe-go/customer"
	"os"
)

type Repository interface {
	CreateCharge(charge *pb.Charge) (*stripe.Charge, error)
	CreateCustomer(charge *pb.Customer) (*stripe.Customer, error)
}

type PaymentRepository struct {
	db *gorm.DB
}

// Create a new user
func (repo *PaymentRepository) CreateTransaction(transaction *pb.Transaction) error {
	if err := repo.db.Create(transaction).Error; err != nil {
		return err
	}
	return nil
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