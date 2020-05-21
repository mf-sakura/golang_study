build:
	docker-compose build

up:
	docker-compose up

create_users:
	curl -X POST localhost:8999/users -H 'Content-Type: application/json' -d '{"first_name": "first", "last_name": "last"}'

show_users:
	curl localhost:8999/users

down:
	docker-compose down
