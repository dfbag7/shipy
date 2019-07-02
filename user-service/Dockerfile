FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-user-service

COPY . .

RUN go mod download

RUN go build -o shipy-user-service main.go repository.go handler.go database.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shipy-user-service/shipy-user-service .

CMD ["./shipy-user-service"]
