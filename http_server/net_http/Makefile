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

incr:
	curl -X POST localhost:8080/incr -d '{"num":3}'
