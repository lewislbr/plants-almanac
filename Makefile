setup:
	@(pre-commit --version && pre-commit install) || (brew install pre-commit && pre-commit install)

start:
	@docker compose -p plantdex up --build -d

start-%:
	@docker compose -p plantdex up --build $*

stop:
	@docker compose -p plantdex down
