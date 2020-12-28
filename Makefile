start:
	@docker-compose -f docker-compose-dev.yaml -p plantdex up -d --build

start-%:
	@docker-compose -f docker-compose-dev.yaml -p plantdex up --build $*

stop:
	@docker-compose -f docker-compose-dev.yaml -p plantdex down
