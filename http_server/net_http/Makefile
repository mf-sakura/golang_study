build:
	go build -o bin/http_server

hello:
	curl localhost:8080/hello

unauthorized:
	curl -i localhost:8080/401

square:
	curl localhost:8080/square -H "num:2"

square_bad_request:
	curl -i localhost:8080/square -H "num:a"

square_with_100:
	curl -i localhost:8080/square -H "num:100"

incr:
	curl -X POST localhost:8080/incr -d '{"num":3}'

incr_by_get:
	curl -X GET localhost:8080/incr -d '{"num":3}'

authn:
	curl -u username:password localhost:8080/auth

authn_invalid:
	curl localhost:8080/auth
