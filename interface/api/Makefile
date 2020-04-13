DBNAME:=golang_study
DATASOURCE:=user:password@tcp(127.0.0.1:3314)/$(DBNAME)?parseTime=true

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

setup: .env install migrate/init migrate/up

mysql:
	mysql -u user -h 127.0.0.1 --protocol tcp --port 3314 -p $(DBNAME)

migrate/init:
	mysql -u user -h 127.0.0.1 --protocol tcp --port 3314 -e "create database if not exists \`$(DBNAME)\` " -p

migrate/status:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" status
migrate/up:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" up

migrate/down:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" down

.env:
	cp $@.example $@

vendor:
	go mod vendor

build:
	go build -o bin/api

create_users:
	curl -X POST localhost:8080/users -H 'Content-Type: application/json' -d '{"first_name": "first", "last_name": "last"}'

show_users:
	curl localhost:8080/users -H 'Content-Type: application/json'

coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
