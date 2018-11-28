build:
	docker build -t service-cli .

run:
	docker run -e MICRO_REGISTRY=mdns service-cli
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml