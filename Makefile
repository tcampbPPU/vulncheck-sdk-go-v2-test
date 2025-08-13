test:
	go test ./...

dbug:
	@go get github.com/dbugapp/dbug-go

format:
	go fmt  ./...

lint:
	@golangci-lint run

# CLI shortcuts for SDK testing
run:
	@if [ -z "$(FUNC)" ]; then \
		echo "Usage: make run FUNC=<function-name>"; \
		echo "Example: make run FUNC=index-vulnrichment"; \
		echo "Use 'make list' to see available functions"; \
	else \
		go run main.go run $(FUNC); \
	fi

list:
	@go run main.go list

help:
	@go run main.go help