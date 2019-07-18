#!/usr/bin/env bash
curl -XPOST -H 'Content-Type: application/json' \
    -d '{"service": "shippy.consignment", "method": "ConsignmentService.Create", "request": { "description": "This is a test", "weight": "500", "containers": [] }}' 
    --url http://localhost:8080/rpc

