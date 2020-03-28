mod:
	go mod vendor
build:
	go build -o bin/echo_server

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
	curl -X POST localhost:8080/incr -H 'Content-Type: application/json' -d '{"num":3}'

incr_method_not_allowed:
	curl -i -X PUT localhost:8080/incr -H 'Content-Type: application/json' -d '{"num":3}'
