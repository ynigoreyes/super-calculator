all: protogen docker-build docker-push

# Generates the protobuffs for every Go service
protogen:
	protoc \
		-I ./infrastructure/proto \
		--go_out=plugins=grpc:./api/Calculator/proto \
		./infrastructure/proto/calculator.proto
	protoc \
		-I ./infrastructure/proto \
		--go_out=plugins=grpc:./api/Client/proto \
		./infrastructure/proto/calculator.proto

docker-build:
	docker build -t miggylol/client ./api/Client
	docker build -t miggylol/calculator ./api/Calculator

docker-push:
	docker push miggylol/client
	docker push miggylol/calculator
