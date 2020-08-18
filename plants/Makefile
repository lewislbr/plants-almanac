include .env

build-dev:
	@docker build -t plants-dev -f Dockerfile.dev ./

build-prod:
	@docker build -t $(PLANTS_IMAGE_NAME) -f Dockerfile.prod ./

run:
	@go run cmd/server/main.go

run-dev:
	@docker run --env-file .env -p 4040:$(PLANTS_APP_PORT) --name plants-dev plants-dev

run-prod:
	@docker run -p 4040:$(PLANTS_APP_PORT) --name plants-prod $(PLANTS_IMAGE_NAME)

stop:
	@docker rm -f $(docker ps -a -q)
