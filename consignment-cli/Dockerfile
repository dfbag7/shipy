FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shipy-consignment-cli

COPY . .

RUN go mod download
#RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o shipy-consignment-cli
RUN go build -o shipy-consignment-cli

FROM alpine:latest

RUN apk --no-cache add ca-certificates

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait
RUN chmod +x /wait

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shipy-consignment-cli/shipy-consignment-cli .
COPY --from=builder /app/shipy-consignment-cli/consignment.json .

CMD /wait && /app/shipy-consignment-cli
