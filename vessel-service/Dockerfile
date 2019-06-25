FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-vessel-service

COPY . .

RUN go mod download
RUN go build -o shipy-vessel-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shipy-vessel-service .

CMD ["./shipy-vessel-service"]
