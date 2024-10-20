.PHONY: build
build:
	go build -v main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run . 

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: migrate
migrate:
	echo "Running migrations up"
	go run ./internal/database/migrate_up.go

.PHONY: migrate_fresh
migrate_fresh:
	echo "Running fresh migrations"
	go run ./internal/database/migrate_fresh.go

.PHONY: migrate_down
migrate_down:
	go run ./migrations/drop_migrations.go

.DEFAULT_GOAL := run