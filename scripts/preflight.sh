#!/usr/bin/env bash
set -euo pipefail
go vet ./...
if command -v golangci-lint >/dev/null 2>&1; then
  golangci-lint run ./...
else
  echo "SKIP: golangci-lint not on PATH (handled by CI's separate lint job)"
fi
go test ./... -count=1 -timeout 120s
go build ./...
