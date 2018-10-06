build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/chat-service \
	    proto/chat/chat.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t chat-service .
run:
	sudo docker run --net="host" --handler=http\
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb\://gcic\:dreamer6@ds163330.mlab.com\:63330/invest \
		chat-service