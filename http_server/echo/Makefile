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

square_large_request:
	curl -i localhost:8080/square -H "num:11"

square_bad_request:
	curl -i localhost:8080/square -H "num:a"

get_incr:
	curl -i localhost:8080/incr

incr:
	curl -X POST localhost:8080/incr -H 'Content-Type: application/json' -d '{"num":3}'

random_hand:
	curl localhost:8080/random_hand
