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

square_bad_request2:
	curl -i localhost:8080/square -H "num:101"

incr:
	curl -X POST localhost:8080/incr -H 'Content-Type: application/json' -d '{"num":3}'

invalid_incr_request:
	curl -i localhost:8080/incr

shared_counter:
	curl -X GET localhost:8080/shared_counter

decr:
	curl -X POST localhost:8080/decr -H 'Content-Type: application/json' -d '{"num":1, "inde_flag":0}'

decr_inde:
	curl -X POST localhost:8080/decr -H 'Content-Type: application/json' -d '{"num":1, "inde_flag":1}'
