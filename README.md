# rss-explore

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/revett/rss-explore)
[![Builds](https://img.shields.io/github/checks-status/revett/rss-explore/main?label=build&style=flat-square)](https://github.com/revett/rss-explore/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/revett/rss-explore?style=flat-square)](https://goreportcard.com/report/github.com/revett/rss-explore)
[![License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/revett/rss-explore/blob/main/LICENSE)

Tools to help you easily find and manage RSS feeds.

## Usage

[**→ 📖 Documentation**](https://revcd.com/projects/rss-explore) (`revcd.com`)

## Examples

Convert a YouTube video URL in to the RSS feed for that video's channel:

```bash
curl -l -s \
  -r POST 'https://rss-explore.revcd.com/youtube/convert' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "url": "https://www.youtube.com/watch?v=7LICrnxWd38"
  }' | jq '.url'
```

```bash
"https://youtube.com/feeds/videos.xml?channel_id=UCWaC58KxHxyCgZVl4cRDxSg"
```

## Development

### Code Generation

```bash
# Install oapi-codegen
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# Generate the API models, server implementation, and spec
make generate
```

### Production

- [Vercel](https://github.com/vercel/vercel) deploys each `.go` file within
  `api/` as a
  [serverless function](https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go)
- Each of these are wrappers around a handler function within the
  `internal/handler` package, which holds the business logic of the API

### Local

```
make run
```

> API available at `http://localhost:5691`

- `main.go` creates a single `echo.Echo` instance, and registers all the handler
  functions from the `internal/handler` package as routes
- This is done via the `api.ServerInterface` that is generated by
  [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen), to make sure that the
  OpenAPI spec is correctly implemented
- This means that get a near-similar experience to production without having to
  mess about with Vercel's `api/` directory structure
