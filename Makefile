build:
	@go build -o bin/minnell cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/minnell