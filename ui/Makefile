build-dev:
	@docker build -t ui-dev -f Dockerfile.dev ./

build-prod:
	@docker build -t ui-prod -f Dockerfile.prod ./

run:
	@npm run start

run-dev:
	@docker run --env-file .env -v ${PWD}:/usr/src/app -p 8080:8080 --name ui-dev ui-dev

run-prod:
	@docker run -p 8080:8080 --name ui-prod ui-prod

stop:
	@docker rm -f $(docker ps -a -q)
