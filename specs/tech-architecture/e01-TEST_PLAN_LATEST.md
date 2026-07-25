# Test Plan — e01 Canary site scaffold

Risk: P1 (pipeline correctness is the point; app logic itself is trivial).

| ID | Scenario | Level | Priority |
|----|----------|-------|----------|
| SC-e01s01-P0-01 | `GET /` returns HTML containing the exact `VERSION` file contents | unit/integration | P0 |
| SC-e01s01-P0-02 | `go vet` and `golangci-lint` pass with zero findings | static | P0 |
| SC-e01s01-P1-01 | `test-build-release.yml` `test` job passes on push to `main` | CI | P1 |
| SC-e01s01-P1-02 | `release` job runs `big-release`, produces tag `v0.1.0` + CHANGELOG entry | CI | P1 |
| SC-e01s01-P1-03 | `deploy.yml` fires via `workflow_run` after Test Build Release succeeds on `main` | CI | P1 |
| SC-e01s01-P1-04 | `bigbase-deploy@v1` step returns HTTP 2xx; health check logs `✅ Site LIVE` | CI | P1 |
| SC-e01s01-P1-05 | `curl https://go.bigbase.click` returns the deployed VERSION in the footer | manual (post-deploy) | P1 |
| SC-e01s01-P2-01 | A second commit produces `v0.1.1` and a fresh successful deploy | manual (regression loop) | P2 |

Fixtures: none — the app has no external dependencies to fixture. `VERSION` file itself is the only test input.
