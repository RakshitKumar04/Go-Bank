build:
	@go build -o bin/Go-Bank

run: build
	@./bin/Go-Bank

test:
	@go test -v ./...