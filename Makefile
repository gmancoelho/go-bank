build:
	@go build -o bin/go-kanban

run: build
	@./bin/go-kanban

test:
	@go test -v ./...