build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/email-service \
	    proto/email/email.proto
	GOOS=linux GOARCH=amd64 go build