testadd:
	curl --request POST localhost:8080/v1/add \
		--header 'content-type: application/json' \
		--data '{"title": "test todo", "priority": 5}'

testget:
	curl --request GET localhost:8080/v1/list

run:
	docker compose up -d

restart:
	docker compose restart

build:
	docker compose build

down:
	docker compose down

logs:
	docker compose logs -f


watch:
	watch docker compose ps

