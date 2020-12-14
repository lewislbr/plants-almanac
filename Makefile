start:
	@docker-compose -f docker-compose-dev.yaml -p plantdex up -d --build

stop:
	@docker-compose -f docker-compose-dev.yaml -p plantdex down
