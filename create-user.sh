curl -XPOST \
	-H 'Content-Type: application/json' \
	-d '{"service": "shipy.auth", "method": "UserService.Create", "request": {"id": "", "email": "ewant.valentine89@gmail.com", "password": "testing123", "name": "Ewan Valentine", "company": "BBC"}}' \
	http://localhost:8080/rpc

