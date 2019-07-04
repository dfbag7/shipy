FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-user-service

COPY . .

RUN go mod download
RUN go build -o shipy-user-service main.go repository.go handler.go database.go token-service.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait
RUN chmod +x /wait

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shipy-user-service/shipy-user-service .

CMD /wait && /app/shipy-user-service
