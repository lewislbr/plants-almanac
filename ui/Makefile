include .env

build:
	@docker build \
	-t plantdex_ui-dev \
	-f Dockerfile.dev \
	./

run:
	@docker run \
	-d \
	-p 8080:$(UI_APP_PORT) \
	-v ${PWD}:/usr/src/app \
	--env-file .env \
	--name plantdex_ui-dev \
	plantdex_ui-dev

stop:
	@docker rm -f $$(docker ps -a -q)
