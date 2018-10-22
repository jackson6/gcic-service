package main

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

func CreateClient(smtpServer SmtpServer) (*smtp.Client, error){
	//build an auth
	auth := smtp.PlainAuth("", smtpServer.addr, smtpServer.password, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.host + ":" + smtpServer.port, tlsconfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return nil, err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}
	return client, nil
}
