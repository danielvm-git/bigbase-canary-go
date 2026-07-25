# Contributing to bigbase-canary-go

Thanks for your interest in contributing! This document explains how to
contribute to **bigbase-canary-go** — a minimal Go HTTP canary site that
serves as a regression signal for the big-release + bigbase-deploy pipeline.

## Project context

This repository is intentionally minimal. Its sole job is to prove, on every
push, that [big-release](https://github.com/danielvm-git/big-release) and the
[bigbase-deploy](https://github.com/danielvm-git/.github) GitHub Action still
work together end-to-end. It is deliberately **not a product** and will never
grow real features.

## How to contribute

1. **Open an issue first** for anything beyond a typo. Describe the gap or
   the problem you're solving.
2. **Fork and branch** from `main`. Use a branch name like
   `fix/typo-readme` or `docs/clarify-architecture`.
3. **Make the change.** Keep it minimal — single `main.go`, no framework,
   no dependencies.
4. **Validate.** Run the preflight: `go vet ./... && golangci-lint run &&
   go test ./... -count=1 -timeout 120s && go build ./...`
5. **Open a pull request** against `main`. Reference the issue it closes.

## Before contributing

Please read:

- [`AGENTS.md`](./AGENTS.md) — project context and agent conventions.
- [`CONVENTIONS.md`](./CONVENTIONS.md) — the quality contract all changes
  must follow.
- [`specs/product/SCOPE_LATEST.yaml`](./specs/product/SCOPE_LATEST.yaml) —
  what's in and out of scope.

## What's welcome

- Bug fixes to the canary itself (the Go server, the CI pipeline).
- Improvements to the release→deploy pipeline integration.
- Documentation that helps other canary maintainers.
- Structural alignment with the big-docs gabarito.

## What's not

- New product features (auth, persistence, multi-page routing, frameworks).
- The canary's value is being boring and stable. If you want to add
  something that makes it more than a one-file HTTP server, it probably
  belongs in a different repo.

## Conventional Commits

All commits must follow [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/).
The CI pipeline checks this, and big-release uses it to determine the next
version.

## Code of conduct

By participating you agree to abide by our [Code of Conduct](./CODE_OF_CONDUCT.md).
