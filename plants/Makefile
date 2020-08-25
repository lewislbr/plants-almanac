include .env

build:
	@docker build \
	-t plantdex_plants-app-dev \
	-f Dockerfile.dev \
	./

run:
	@[ -d plants-db ] || mkdir plants-db

	@docker network create plants

	@docker run \
	-d \
	--network plants \
	-v ${PWD}/plants-db:/data/db \
	--name plantdex_plants-db-dev \
	mongo:4

	@docker run \
	-d \
	-p 127.0.0.1:4040:${PLANTS_APP_PORT} \
	--network plants \
	--env-file .env \
	--name plantdex_plants-app-dev \
	plantdex_plants-app-dev

stop:
	@docker rm -f $$(docker ps -a -q)

	@docker volume prune -f

	@docker network prune -f
