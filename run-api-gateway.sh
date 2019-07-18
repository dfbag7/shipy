#!/usr/bin/env bash
docker run --network=go-ser_default  \
   -p 8080:8080 \
   -e MICRO_REGISTRY=mdns \
   microhq/micro api \
   --handler=rpc \
   --address=:8080 \
   --namespace=shipy \
   --cors-allowed-methods=GET,POST,OPTIONS,HEAD \
   --cors-allowed-origins=*

# The CORS parameters seems not working with the default microhq/micro image,
# because it is not built with the CORS plugin!!!

#   -e CORS_ALLOWED_HEADERS=* \
#   -e CORS_ALLOWED_ORIGINS=* \
#   -e CORS_ALLOWED_METHODS=GET,POST,OPTIONS,HEAD \
