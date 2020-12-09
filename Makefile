start:
	@docker-compose -f docker-compose-dev.yaml -p plantdex up -d --build

stop:
	@docker-compose -f docker-compose-dev.yaml -p plantdex down && \
	docker volume rm $$(docker volume ls -q | awk -F, 'length($0) == 64 { print }')
