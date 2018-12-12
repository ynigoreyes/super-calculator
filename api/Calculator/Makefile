# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=sandbox
BINARY_UNIX=$(BINARY_NAME)_unix

all: say-hello

say-hello:
	echo "Make works"

dev:
	$(GOCMD) run main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

protogen:
	protoc -I ../../infrastructure ../../infrastructure/proto/calculator.proto --go_out=plugins=grpc:../../api/Calculator/proto

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker build -t miggylol/calculator .
