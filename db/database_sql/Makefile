DBNAME:=golang_study
DATASOURCE:=root:password@tcp(127.0.0.1:3306)/$(DBNAME)

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

setup: install docker-compose/up migrate/init migrate/up

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DBNAME)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/status:
	goose status

migrate/up:
	goose up

migrate/down:
	goose down

.env:
	cp $@.example $@
