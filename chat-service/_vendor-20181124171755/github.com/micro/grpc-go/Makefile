all: vet test testrace testappengine

build: deps
	go build github.com/micro/grpc-go/...

clean:
	go clean -i github.com/micro/grpc-go/...

deps:
	go get -d -v github.com/micro/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/micro/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 5m github.com/micro/grpc-go/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 5m github.com/micro/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/micro/grpc-go/...

testdeps:
	go get -d -v -t github.com/micro/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/micro/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/micro/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/micro/grpc-go/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
