# Conventions — bigbase-canary-go

Governed by the [bigpowers](https://github.com/danielvm-git/bigpowers) methodology. This file is the project-local subset relevant to a minimal single-file canary — see bigpowers `CONVENTIONS.md` for the full doctrine.

## Conventional Commits & Semantic Versioning

All changes MUST follow [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/). Versioning follows [Semantic Versioning 2.0.0](https://semver.org/), decided at release time by `big-release` from commit history — never hand-tracked.

**Format:** `<type>(<scope>): <description>` (space after colon mandatory).

- `feat`: Minor bump — new feature
- `fix`: Patch bump — bug fix
- `perf`: Patch bump — performance improvement
- `docs`, `chore`, `style`, `refactor`, `test`: No bump (unless breaking)
- `BREAKING CHANGE:` (or `!` after type): Major bump

## GitHub & Git Operations

- No direct work on `main`. Every task starts with a feature branch/worktree via `kickoff-branch`.
- Integrate (solo-git profile): `bash scripts/land-branch.sh <branch> "<conventional message>"` after `release-branch` gates — local squash to `main`, then push.
- Use `gh repo clone`, not `git clone`. Use `gh run watch` / `gh pr checks` for CI status.
- **Git Attribution:** NEVER include `Co-authored-by:` or any AI-agent attribution footer.
- Never call GitHub REST API directly (curl/fetch) — use `gh`. (bigbase's own REST API is a separate service and is fine to curl directly for provisioning.)
- Never create GitHub issues from automated workflows — produce local `.md` files in `specs/bugs/` instead.

## Always Green / Shift Left

Preflight and CI MUST be green before any forward work — not "green enough for this task." Fixing a red gate now is cheaper than debugging it in production (1-10-100 rule).

**Preflight:** `go vet ./... && golangci-lint run && go test ./... -count=1 -timeout 120s && go build ./...` — must pass before kickoff, develop, or verify phases advance.

**CI green:** `gh pr checks` (or the Actions tab) must show passing before merge/land.

## Discovered Defects

Any reproducible gate failure found during unrelated work is a discovered defect, not optional noise.

1. **quick-fix** — trivial, single-file fixes within guardrails.
2. **fix-bug** — needs investigation (`specs/bugs/BUG-*.md` + TDD).
3. **Log** — only when reproduction is blocked after a good-faith attempt.

Discovered fixes ship in the same land as the original work, in separate commits.

**Banned dismissive phrases:** "pre-existing", "unrelated to this session", "not introduced by my changes", "out of scope" (when ignoring a red gate) — none of these waive a red gate. Fix or log instead.

## specs/ — All Planning Output Goes Here

- `specs/state.yaml` — active session, handoff
- `specs/release-plan.yaml` — release index, story BCPs
- `specs/execution-status.yaml` — sole source of truth for story done/pending status
- `specs/epics/e01-canary-site/` — this repo's one epic
- `specs/bugs/BUG-*.md` + `specs/bugs/registry.yaml` — bug investigations (never GitHub issues)
- `specs/tech-architecture/tech-stack.md` — stack notes

## Defensive Code Categories

None apply. This is a static version-footer endpoint with no external dependencies, no untrusted input beyond the HTTP request line, and no failure modes worth guarding — no rate limiting, retry, circuit breaker, timeout, or graceful-degradation logic. Adding any would be complexity the app doesn't need.

## Risk Tier

P2 — infrastructure/regression-signal repo, not user-facing product. `plan-tests` still runs (the release→deploy pipeline's correctness is the actual point) but `test_plan` is not waived; BCP sizing stays lightweight.
