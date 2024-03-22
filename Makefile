all:
	make tidy
	make test
	make lint

tidy:
	go mod tidy
	go vet ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54
	golangci-lint run ./...

test:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm -f coverage.out

test_short:
	go test -short ./...

clean:
	go clean