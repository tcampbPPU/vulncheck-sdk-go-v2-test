test:
	go test ./...

dbug:
	@go get github.com/dbugapp/dbug-go

format:
	go fmt  ./...

lint:
	@golangci-lint run