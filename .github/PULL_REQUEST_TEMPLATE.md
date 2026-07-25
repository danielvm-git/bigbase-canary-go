<!--
  Thanks for contributing to bigbase-canary-go! Fill in the sections below.
-->

## Summary

<!-- One or two sentences: what does this PR change, and why? -->

## Linked issue

<!-- "Closes #N" or "Refs #N". -->

## What kind of change is this?

- [ ] Bug fix (canary server or pipeline)
- [ ] Pipeline / CI improvement
- [ ] Documentation or specs
- [ ] Structural alignment (big-docs gabarito)
- [ ] Other

## Checklist

- [ ] I have read [`CONTRIBUTING.md`](../CONTRIBUTING.md) and
  [`CONVENTIONS.md`](../CONVENTIONS.md).
- [ ] The change keeps the server minimal (single `main.go`, no framework).
- [ ] No external dependencies were added.
- [ ] Preflight passes: `go vet ./... && go test ./... -count=1 -timeout 120s && go build ./...`
- [ ] Conventional Commits on every commit.
- [ ] One trailing newline at the end of every file.
