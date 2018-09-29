package main

import "net/smtp"

// CreateClient creates the main client to our email instance
func CreateClient(addr string) (*smtp.Client, error) {
	client, err := smtp.Dial(addr)
	if err != nil {
		return nil, err
	}

	return client, nil
}
