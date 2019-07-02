FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-user-cli

COPY . .

RUN go mod download

RUN go build -o shipy-user-cli

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shipy-user-cli/shipy-user-cli .

CMD ["./shipy-user-cli"]
