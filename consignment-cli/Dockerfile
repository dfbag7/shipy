FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-consignment-cli

COPY . .

RUN go mod download

#RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o shipy-consignment-cli
RUN go build -o shipy-consignment-cli

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shipy-consignment-cli/shipy-consignment-cli .
COPY --from=builder /app/shipy-consignment-cli/consignment.json .

CMD ["./shipy-consignment-cli"]
