# Tech Stack — bigbase-canary-go

- **Language/runtime:** Go 1.22+, standard library only (`net/http`).
- **Modules:** single `main.go` — one handler, reads `VERSION`, writes an HTML footer.
- **CI:** GitHub Actions, `.github/workflows/test-build-release.yml` (lint → test → build → release) + `.github/workflows/deploy.yml` (triggered via `workflow_run`), copied from `danielvm-git/.github`'s `test-build-release-go.yml`/`deploy-go.yml` templates.
- **Release:** [big-release](https://github.com/danielvm-git/big-release) (`.big-release.yml`: `tagFormat: v${version}`, plugins `changelog`/`git`/`github`), replacing the template's default `semantic-release` step — this repo exists specifically to exercise big-release, not semantic-release.
- **Deploy:** `danielvm-git/.github/actions/bigbase-deploy@v1` → bigbase site `go` → `https://go.bigbase.click`, `app_type: go`.
- **Gray areas:** none load-bearing — no error handling beyond what `net/http` gives for free; no config beyond `VERSION`.
