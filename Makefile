build:
	@go build -o bin/postgres_go_rest_api

run: build
	@./bin/postgres_go_rest_api

test:
	@go test -v ./...