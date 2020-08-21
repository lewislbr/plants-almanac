include .env

build:
	@docker build \
	-t plants-almanac_plants-app-dev \
	-f Dockerfile.dev \
	./

run:
	@docker volume create \
	--label keep=true \
	plants-almanac_plants-db

	@docker run \
	-d \
	-p 27017:27017 \
	-v plants-almanac_plants-db:/data/db \
	--name plants-almanac_plants-db-dev \
	mongo:4

	@docker run \
	-d \
	-p 4040:$(PLANTS_APP_PORT) \
	--env-file .env \
	--name plants-almanac_plants-app-dev \
	plants-almanac_plants-app-dev

stop:
	@docker rm -f $$(docker ps -a -q)

	@docker volume prune -f --filter "label!=keep"
