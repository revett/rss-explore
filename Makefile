generate:
	@go generate ./...

lint:
	@golangci-lint run

run:
	@go build .
	@./rss-explore

test:
	@go test -cover ./...
