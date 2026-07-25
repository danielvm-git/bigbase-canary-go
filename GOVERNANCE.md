# Governance

## Project status

**bigbase-canary-go** is a minimal Go HTTP canary site that exists solely as a
regression signal for the big-release + bigbase-deploy pipeline. It is
deliberately small, boring, and stable. No new product features will be added.

## Decision-making structure

| Role | Who | Decides |
| --- | --- | --- |
| **Maintainer** | @danielvm-git | Final say on all decisions, what gets merged, and project direction. |
| **Contributors** | Anyone who opens a merged PR | Propose changes via issues and pull requests. |
| **Community** | Anyone reading or using the canary | Raises questions in Discussions; influences direction through feedback. |

## How decisions are made

1. **Small, obvious fixes** (typos, CI pin updates, dependency bumps) — a
   contributor opens a PR; the maintainer reviews and merges. No ceremony.

2. **Changes to the canary itself** (the Go server, the VERSION mechanism,
   the health check) — open an issue first. The change must keep the server
   minimal: single `main.go`, no framework, no external dependencies.

3. **Changes to the CI/CD pipeline** (workflows, big-release configuration,
   bigbase-deploy integration) — these are high-stakes because the pipeline
   *is* the product. Changes require an issue, a proposal, and a successful
   end-to-end test.

4. **Changes to project scope** — adding anything beyond a one-file HTTP
   server requires changing `specs/product/SCOPE_LATEST.yaml` first. Most
   additions will be rejected — this repo's value is being minimal.

## Solo-git workflow

This project uses a solo-git integration profile:

- All development happens on feature branches created via `kickoff-branch`.
- Integration uses `bash scripts/land-branch.sh` — local squash to `main`,
  then push.
- No pull request ceremony required for solo work, but PRs are welcome for
  community contributions.
- `big-release` cuts semver tags from Conventional Commits on every push.

## The team

Maintainers and regular contributors are listed in [`CODEOWNERS`](./CODEOWNERS).

## Amendments to this document

Propose changes by opening a PR that updates this file. Explain the governance
gap the change addresses.
