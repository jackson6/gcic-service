build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/email-service \
	    proto/email/email.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t email-service .
run:
	sudo docker run --net="host" \
		-p 50051 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e DISABLE_AUTH=false \
		-e MICRO_REGISTRY=mdns \
		email-service
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml