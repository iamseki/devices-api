DATABASE_URI ?= postgres://devices:devices@localhost:5432/devices?sslmode=disable
MIGRATIONS_DIR := migrations

.PHONY: run-migrations rollback-migration force-migration help build lint test clean sql-command

run-migrations:
	migrate -database ${DATABASE_URI} -path ${MIGRATIONS_DIR} up

rollback-migration:
	migrate -database ${DATABASE_URI} -path ${MIGRATIONS_DIR} down 1

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a migration name using 'name=<migration_name>'"; \
		exit 1; \
	fi
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $(name)

lint:
	golangci-lint run ./...

test:
	go test ./src/domain -v

generate-code-sql:
	sqlc generate

sql-command:
	psql ${DATABASE_URI} --command="$(sql)"

help:
	@echo "Available targets:"
	@echo "  run-migrations    Apply all pending migrations"
	@echo "  rollback-migration Rollback the last applied migration"
	@echo "  create-migration  Create a new migration (use name=<migration_name>)"
	@echo "  lint              Run code linting"
	@echo "  test              Run tests"
	@echo "  generate-code-sql Compile SQL to type-safe golang code with sqlc"
	@echo "  sql-command  Execute a SQL command to DATABASE_URI (use sql='<sql_command>')"