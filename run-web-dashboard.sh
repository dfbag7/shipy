docker run --network=go-ser_default  -p 8080:8080 -e MICRO_REGISTRY=mdns microhq/micro web --address=:8080 --namespace=shipy

