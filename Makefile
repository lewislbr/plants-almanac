start:
	@docker compose -p plantdex up --build -d

start-%:
	@docker compose -p plantdex up --build $*

stop:
	@docker compose -p plantdex down
