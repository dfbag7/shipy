build:
	protoc -I. --go_out=plugins=micro:. \
		proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build

container: build
	docker build -t shipy-vessel-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shipy-vessel-service
