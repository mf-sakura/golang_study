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

create_user:
	curl -X POST localhost:8080/users -H 'Content-Type: application/json' -d '{"first_name":"折木", "last_name":"奉太郎"}'

.env:
	cp $@.example $@

vendor:
	go mod vendor

build:
	go build -o bin/sqlx
