build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build main.go repository.go handler.go datastore.go

container: build
	docker build -t shipy-consignment-service .

run:
	docker run -p 50051:50051 \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    shipy-consignment-service

