.PHONY: test linter
test:
	go test -v ./... $(TEST_ARGS)

linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --timeout=5m -E ginkgolinter
