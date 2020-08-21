include .env

build:
	@docker build \
	-t plants-almanac_ui-dev \
	-f Dockerfile.dev \
	./

run:
	@docker run \
	-d \
	-p 8080:$(UI_APP_PORT) \
	-v ${PWD}:/usr/src/app \
	--env-file .env \
	--name plants-almanac_ui-dev \
	plants-almanac_ui-dev

stop:
	@docker rm -f $$(docker ps -a -q)
