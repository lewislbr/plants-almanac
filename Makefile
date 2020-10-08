start:
	@docker-compose -f docker-compose-dev.yml up -d --build

stop:
	@docker-compose -f docker-compose-dev.yml down

	@docker volume rm $$(docker volume ls -q | awk -F, 'length($0) == 64 { print }')
