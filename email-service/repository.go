package main

import (
	"log"
	"net/smtp"
)

type Repository interface {
	Send(*Mail) error
	Close()
}

type EmailRepository struct {
	client *smtp.Client
}

func (repo *EmailRepository) Send(message *Mail) error {
	// step 2: add all from and to
	if err := repo.client.Mail(message.to[0]); err != nil {
		log.Panic(err)
	}
	receivers := append(message.to)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err := repo.client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := repo.client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message.message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	repo.Close()

	return nil
}

func (repo *EmailRepository) Close() {
	repo.client.Quit()
}