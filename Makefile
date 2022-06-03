.PHONY: build
build:
	docker-compose build

up:
	@make build
	docker-compose up
