# e01s01 — Stand up the Go canary site end-to-end

## Story

As the bigbase platform maintainer, I want a minimal Go site deployed through the real
big-release + bigbase-deploy pipeline, so that any future change to either tool has an
immediate, low-noise regression signal.

## Acceptance criteria

1. `GET /` on the running app returns HTML containing the exact contents of `VERSION`.
2. `.github/workflows/test-build-release.yml` lint/test/build jobs pass on a push to `main`.
3. The `release` job runs `big-release release --verbose` and produces a real git tag (`v0.1.0`) + `CHANGELOG.md` entry from Conventional Commits — not `semantic-release`.
4. `.github/workflows/deploy.yml` fires via `workflow_run` after a successful Test Build Release run on `main`, and its `bigbase-deploy@v1` step returns HTTP 2xx from `POST /api/sites/{id}/deploy`.
5. The deploy job's health-check step logs `✅ Site LIVE`.
6. `curl https://go.bigbase.click` returns the footer with the deployed `VERSION`.
7. A second Conventional Commit (e.g. `fix: adjust footer copy`) produces `v0.1.1` and a fresh deploy — proving the loop, not just first boot.

## Out of scope

Anything in `specs/product/SCOPE_LATEST.yaml`'s `out_of_scope` list.
