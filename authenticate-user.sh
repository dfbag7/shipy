curl -XPOST -H 'Content-Type: application/json' \ 
    -d '{ "service": "shippy.auth", "method": "Auth.Auth", "request":  { "email": "your@email.com", "password": "SomePass" } }' \
	        http://localhost:8080/rpc

