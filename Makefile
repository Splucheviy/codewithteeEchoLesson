.PHONY: build
build:
	go build -v cmd/api/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: tidy
tidy:
	go mod tidy
.DEFAULT_GOAL := run