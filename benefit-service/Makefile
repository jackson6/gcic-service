build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/benefit-service \
	    proto/benefit/benefit.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t benefit-service .
run:
	sudo docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50056 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb\://gcic\:dreamer6@ds163330.mlab.com\:63330/invest \
		benefit-service