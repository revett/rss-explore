# rss-explore

Tools to help you easily find and manage RSS feeds.

## Production

- [Vercel](https://github.com/vercel/vercel) deploys each `.go` file within
  `api/` as a
  [serverless function](https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go)
- Each of these are wrappers around a handler function within the
  `internal/handler` package, which holds the business logic of the API

## Local Development

```
make run
```

> API available at `http://localhost:5691`

- `main.go` creates a single `echo.Echo` instance, and registers all the handler
  functions from the `internal/handler` package as routes
- This means that get a near-similar experience to production without having to
  mess about with Vercel's `api/` directory structure
