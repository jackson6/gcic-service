package main

import (
	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	db, err := CreateSession()


	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer db.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore - %v", err)
	}
	db.AutoMigrate(&paymentService.Transaction{})

	repo := &PaymentRepository{db}

	srv := micro.NewService(
		micro.Name("gcic.payment"),
		micro.Version("latest"),
	)

	srv.Init()

	paymentService.RegisterPaymentServiceHandler(srv.Server(), &service{ repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
