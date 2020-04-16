setup:
	docker-compose up -d

build:
	go build -o bin/interface_example
create_bucket:
	go run script/create_bucket/main.go

list_buckets:
	go run script/list_buckets/main.go

list_objects:
	aws s3 --endpoint-url http://localhost:4572 ls test/output/

s3_write:
	bin/interface_example -s test.txt -a write -p aws
s3_read:
	bin/interface_example -s test.txt -a read -p aws

local_write:
	bin/interface_example -s test.txt -a write -p local

local_read:
	bin/interface_example -s test.txt -a read -p local

coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
