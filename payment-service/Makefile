build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/payment-service \
	    proto/payment/payment.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t payment-service .
run:
	sudo docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		-e STRIPE_KEY="sk_test_0R7KtTokmbmUk1sK2qqdyi6y"
		-e DB_HOST=mongodb\://gcic\:dreamer6@ds163330.mlab.com\:63330/invest \
		payment-service