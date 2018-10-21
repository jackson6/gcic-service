package main

import "time"

type Transaction struct {
	Class string `json:"$class"`
	Points float64 `json:"points"`
	Member string `json:"member"`
	Partner string `json:"partner"`
	TransactionId string `json:"transactionId"`
	Timestamp time.Time `json:"timestamp"`
}

type Member struct {
	Class string `json:"$class"`
	MemberId string `json:"memberId"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	CardNumber string `json:"cardNumber"`
	Points float64 `json:"points"`
}

type Partner struct {
	Class string `json:"$class"`
	PartnerId string `json:"partnerId"`
	Name string `json:"name"`
	Address string `json:"address"`
}