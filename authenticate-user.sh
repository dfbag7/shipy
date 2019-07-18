#!/usr/bin/env bash
curl -XPOST -H 'Content-Type: application/json' \
    -d '{ "service": "shipy.auth", "method": "UserService.Auth", "request":  { "email": "ewant.valentine89@gmail.com", "password": "testing123" } }' \
	        http://localhost:8080/rpc
