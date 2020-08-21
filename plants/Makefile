include .env

build:
	@docker build \
	-t plants-almanac_plants-app-dev \
	-f Dockerfile.dev \
	./

run:
	@docker run \
	-d \
	-p 4040:$(PLANTS_APP_PORT) \
	--env-file .env \
	--name plants-almanac_plants-app-dev \
	plants-almanac_plants-app-dev

stop:
	@docker rm -f $$(docker ps -a -q)
