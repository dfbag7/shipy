build:
	protoc -I. --go_out=plugins=micro:. \
		proto/auth/auth.proto
	GOOS=linux GOARCH=amd64 go build

