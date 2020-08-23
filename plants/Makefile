include .env

build:
	@docker build \
	-t plantdex_plants-app-dev \
	-f Dockerfile.dev \
	./

run:
	@[ -d plants-db ] || mkdir plants-db

	@docker run \
	-d \
	-p 27017:27017 \
	-v ${PWD}/plants-db:/data/db \
	--name plantdex_plants-db-dev \
	mongo:4

	@docker run \
	-d \
	-p 4040:${PLANTS_APP_PORT} \
	--env-file .env \
	--name plantdex_plants-app-dev \
	plantdex_plants-app-dev

stop:
	@docker rm -f $$(docker ps -a -q)

	@docker volume prune -f
