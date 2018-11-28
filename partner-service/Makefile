build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/partner-service \
	    proto/partner/partner.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t partner-service .
run:
	sudo docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50055 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb\://gcic\:dreamer6@ds163330.mlab.com\:63330/invest \
		partner-service
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml