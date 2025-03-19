
- migrate cli install with go: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` to manage migrations in the database
- sqlc install with go: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` to generate type safe code from sql queries
- go install github.com/swaggo/swag/cmd/swag@latest 
  - swag init
- make run-migrations


// TODO:
- integration test []
  - including specific use cases []
- refactor []
- k6
- indexes
- telemetry
- readme