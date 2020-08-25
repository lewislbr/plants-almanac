include .env

build:
	@docker build \
	-t plantdex_ui-app-dev \
	-f Dockerfile.dev \
	./

run:
	@docker run \
	-d \
	-p 127.0.0.1:8080:${UI_APP_PORT} \
	-v ${PWD}:/usr/src/app \
	--env-file .env \
	--name plantdex_ui-app-dev \
	plantdex_ui-app-dev

stop:
	@docker rm -f $$(docker ps -a -q)
