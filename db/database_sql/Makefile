DBNAME:=golang_study
DATASOURCE:=root:password@tcp(localhost:3306)/$(DBNAME)

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

setup: .env install docker-compose/up migrate/init migrate/up

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DBNAME)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/status:
	goose -dir db/migrations mysql "$(DATASOURCE)" status

migrate/up:
	goose -dir db/migrations mysql "$(DATASOURCE)" up

migrate/down:
	goose -dir db/migrations mysql "$(DATASOURCE)" down

.env:
	cp $@.example $@
