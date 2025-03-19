A simple API for managing devices, written in Golang, featuring CRUD endpoints and strict business rules for updating and deleting devices.


The API was tested and developmed in the following environment:

- *Docker Version:* Docker Engine - Community 20.10.22
- *Golang Version:* go1.24.0 - linux/amd64


- migrate cli install with go: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` to manage migrations in the database
- sqlc install with go: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` to generate type safe code from sql queries
- go install github.com/swaggo/swag/cmd/swag@latest 
  - swag init
- make run-migrations


## Setup

- To run the application and the database execute: `docker compose up`
- Then, the API documentation should be acessed through: http://localhost:8081/swagger/index.html

## Design Decisions

### Domain centric

### sqlc

### migrate

### Echo Http Framework

### API documentation

- http://localhost:8081/swagger/index.html

## Setup

## Index Creation on Brand

## Running domain tests

- `go test ./src/domain -v` or `make test`

// TODO:
- boilerplate + sqlc + migrations [x]
- echo router + swagger [x]
- refactor and test adding domain layer, domain centric architecture, side effects on the edges [x]

improvements:
- integration test []
  - including specific use cases []
- open telemetry []
- refactor []
  - domain centric architecture? side effect on edges
- k6
- indexes

- readme