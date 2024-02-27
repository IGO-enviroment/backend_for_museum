include .env
export

compose-up: ### Run docker-compose
	cd ./infra && docker-compose up --build -d postgres redis && docker-compose logs -f

swag-v1: ### swag init
	swag init -g config/web/v1/routes.go

run: swag-v1
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' CGO_ENABLED=0 go run -tags migrate ./cmd

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations '$(MIGRATE_NAME)'

migrate-up:
	migrate -path migrations/ -database '$(PG_URL)?sslmode=disable' up

migrate-down:
	migrate -path migrations/ -database '$(PG_URL)?sslmode=disable' down -all

linter-golangci: ### check by golangci linter
	golangci-lint run
