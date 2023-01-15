generate:
	@oapi-codegen --config oapi-codegen.yml rss-explore.openapi.yml

lint:
	@golangci-lint run

run: generate
	@go build .
	@./rss-explore

test:
	@go test -cover ./...
