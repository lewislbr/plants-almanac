build-dev:
	@docker build -t plants-dev -f Dockerfile.dev ./

build-prod:
	@docker build -t plants-prod -f Dockerfile.prod ./

run:
	@go run cmd/server/main.go

run-dev:
	@docker run --env-file .env -p 4040:4040 --name plants-dev plants-dev

run-prod:
	@docker run -p 4040:4040 --name plants-prod plants-prod

stop:
	@docker rm -f $(docker ps -a -q)
