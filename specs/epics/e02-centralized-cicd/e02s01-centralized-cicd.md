# Story e02s01: Wire central CI/CD reusable workflow and decommission legacy workflows

**type:** chore  
**risk:** P1  
**context:** infra  
**BCPs:** 2  

## Context
Adapt `bigbase-canary-go` to the centralized v4.1.0 portfolio CI/CD architecture described in `danielvm-git/.github` Discussion #32. Replace standalone `.github/workflows/test-build-release.yml` and `.github/workflows/deploy.yml` with a single `.github/workflows/ci-cd.yml` that inherits `danielvm-git/.github/.github/workflows/test-build-release-go.yml@main` passing `site_url: "https://go.bigbase.click"` and `release_tool: "big-release"`.

## Requirements Delta

#### MODIFIED: CI/CD Pipeline Execution
**Before:** CI/CD executed via local standalone workflows (`test-build-release.yml` and `deploy.yml`) maintained inside `bigbase-canary-go`.  
**After:** CI/CD executed via reusable workflow `danielvm-git/.github/.github/workflows/test-build-release-go.yml@main` with parameters `site_url: "https://go.bigbase.click"` and `release_tool: "big-release"`.

#### REMOVED: Standalone Workflow Duplication
**Before:** Duplicate copy of linting, testing, building, and deploying workflow logic in `.github/workflows/test-build-release.yml` and `.github/workflows/deploy.yml`.  
**After:** (removed) — obsolete files decommissioned; single `.github/workflows/ci-cd.yml` entry point.

## Steps

1. Create `.github/workflows/ci-cd.yml` calling `uses: danielvm-git/.github/.github/workflows/test-build-release-go.yml@main` with `with: site_url: "https://go.bigbase.click"` and `release_tool: "big-release"` → verify: `test -f .github/workflows/ci-cd.yml && grep -q 'uses: danielvm-git/.github/.github/workflows/test-build-release-go.yml@main' .github/workflows/ci-cd.yml`
2. Decommission legacy `.github/workflows/test-build-release.yml` and `.github/workflows/deploy.yml` → verify: `[ ! -f .github/workflows/test-build-release.yml ] && [ ! -f .github/workflows/deploy.yml ]`
3. Execute preflight validation suite → verify: `go vet ./... && golangci-lint run && go test ./... -count=1 -timeout 120s && go build ./...`

## Verification Script (Step-by-Step)

1. Inspect `.github/workflows/ci-cd.yml` to confirm reusable workflow reference and parameter configuration.
2. Confirm `.github/workflows/` contains only `ci-cd.yml`.
3. Run `go vet ./... && golangci-lint run && go test ./... -count=1 -timeout 120s && go build ./...` to confirm local environment remains green.

## Out of scope
- Application standard library server code (`main.go`).
- Modifying `big-release` binary functionality.

## Risks
- Incorrect reusable workflow path syntax: Mitigated by checking `danielvm-git/.github` repository structure.
