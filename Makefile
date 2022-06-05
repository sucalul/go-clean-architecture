.PHONY: build
build:
	docker-compose build

up:
	@make build
	docker-compose up

exec:
	docker-compose exec api bash

exec-db:
	docker-compose exec db bash -c "psql -U postgres"

