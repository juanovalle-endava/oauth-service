lint:
	golangci-lint run -v
test:
	go clean -testcache && go test -v -cover ./...
run:
	go run cmd/main.go