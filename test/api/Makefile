DBNAME:=golang_study
TEST_DBNAME=golang_study_test
DATASOURCE:=root:rootpassword@tcp(127.0.0.1:3314)/$(DBNAME)?parseTime=true
TEST_DATASOURCE:=root:rootpassword@tcp(127.0.0.1:3314)/$(TEST_DBNAME)?parseTime=true

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

setup: .env install migrate/init migrate/up

mysql:
	mysql -u root -h 127.0.0.1 --protocol tcp --port 3314 -prootpassword $(DBNAME)

mysql/test:
	mysql -u root -h 127.0.0.1 --protocol tcp --port 3314 -prootpassword $(TEST_DBNAME)

migrate/init:
	mysql -u root -h 127.0.0.1 --protocol tcp --port 3314 -e "create database if not exists \`$(DBNAME)\` " -prootpassword
	mysql -u root -h 127.0.0.1 --protocol tcp --port 3314 -e "create database if not exists \`$(TEST_DBNAME)\` " -prootpassword

migrate/status:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" status

migrate/up:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" up

migrate/down:
	goose -dir "db/migrations" mysql "$(DATASOURCE)" down

migrate/status/test:
	goose -dir "db/migrations" mysql "$(TEST_DATASOURCE)" status

migrate/up/test:
	goose -dir "db/migrations" mysql "$(TEST_DATASOURCE)" up

migrate/down/test:
	goose -dir "db/migrations" mysql "$(TEST_DATASOURCE)" down

.env:
	cp $@.example $@

vendor:
	go mod vendor

build:
	go build -o bin/test

create_users:
	curl -X POST localhost:8080/users -H 'Content-Type: application/json' -d '{"first_name": "first", "last_name": "last"}'

show_users:
	curl localhost:8080/users

test:
	go test ./...

coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
