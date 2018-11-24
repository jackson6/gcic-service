build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/loyalty-serv \
	    proto/loyalty/loyalty.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t loyalty-ser .
run:
	docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		loyalty-serv