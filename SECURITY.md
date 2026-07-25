# Security Policy

## Supported Versions

Only the latest release is supported. This is a canary/regression-signal
project — older versions are not maintained.

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |
| < latest | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. Thank you for improving the security
of this project.

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, report them privately using one of these methods:

- **GitHub private vulnerability reporting** — use the
  *"Report a vulnerability"* button under the **Security** tab of this
  repository (preferred).
- **Email** — send details to **danielvm.gonzalez@gmail.com**.

Please include:

1. A description of the vulnerability and its impact.
2. Steps to reproduce, including any proof-of-concept.
3. Affected versions (if known).
4. Any suggested mitigations.

### Response timeline

- We will acknowledge your report within **48 hours**.
- We will provide an initial assessment and a planned fix date within **7 days**.
- We will notify you when the vulnerability is fixed.

### Scope

In scope: vulnerabilities in this project's own code (the Go server).
Out of scope: vulnerabilities in upstream dependencies (report those upstream),
social engineering, or denial-of-service against the project's infrastructure.

## Disclosure policy

We follow coordinated disclosure: details are published only after a fix is
available and affected users have had reasonable time to upgrade.

## Related security artifacts

This file is the vulnerability-reporting policy. The project's security design
and audit artifacts live under `specs/`:

- `specs/tech-architecture/security-plan.md` — forward-looking defensive design.
- `specs/tech-architecture/impact.md` — impact analysis for changes.
