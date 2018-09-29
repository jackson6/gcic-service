build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/user-service \
	    proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t user-service .
run:
	sudo docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb\://gcic\:dreamer6@ds163330.mlab.com\:63330/invest \
		user-service