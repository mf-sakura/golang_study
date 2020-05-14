PORT=8000

mod:
	go mod vendor

build_local:
	go build -o bin/echo_server

up_local:
	PORT=${PORT} ./bin/echo_server

build_docker:
	docker build -t myapp:latest --build-arg=REVISION=`git rev-parse HEAD` .

up_docker:
	docker run -p ${PORT}:${PORT} -e PORT=${PORT} myapp:latest

up_docker_deamon:
	docker run -d -p ${PORT}:${PORT} -e PORT=${PORT} myapp:latest

hello:
	curl localhost:${PORT}/hello
