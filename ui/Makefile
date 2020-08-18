include .env

build-dev:
	@docker build -t ui-dev -f Dockerfile.dev ./

build-prod:
	@docker build -t $(UI_IMAGE_NAME) -f Dockerfile.prod ./

run:
	@npm run start

run-dev:
	@docker run --env-file .env -v ${PWD}:/usr/src/app -p 8080:$(UI_APP_PORT) --name ui-dev ui-dev

run-prod:
	@docker run -p 8080:$(UI_APP_PORT) --name ui-prod $(UI_IMAGE_NAME)

stop:
	@docker rm -f $(docker ps -a -q)
