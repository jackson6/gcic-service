build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jackson6/gcic-service/websocket-server \
	    proto/chat/chat.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t websocket-server .
run:
	sudo docker run --net="host" chat-service
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml