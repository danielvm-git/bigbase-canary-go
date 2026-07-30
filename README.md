# bigbase-canary-go

A minimal Go HTTP canary site that verifies the [bigbase-deploy](https://github.com/danielvm-git/.github) GitHub Action and [big-release](https://github.com/danielvm-git/big-release) work together end-to-end.

**Live:** https://go.bigbase.click

## Quick Start

```bash
go run .
```

## Test

```bash
go test ./... -count=1 -timeout 120s
```

## How It Works

A single `main.go` serves an HTTP response with a version footer read from the `VERSION` file at request time. CI runs lint, test, build, and release via `big-release`, then deploys through `bigbase-deploy`. The health check confirms the deployed site returns HTTP 200.
