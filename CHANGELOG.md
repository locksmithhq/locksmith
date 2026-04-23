# Changelog

All notable changes to Locksmith are documented here.
Each version has its own detailed file under [`docs/changelog/`](docs/changelog/).

Breaking changes across all versions are aggregated in [`docs/breaking-changes.md`](docs/breaking-changes.md).
Deployment instructions (Docker Compose and `docker run`) are in [`docs/deployment.md`](docs/deployment.md).

---

## How to contribute

- **Ongoing work**: add entries to [`docs/changelog/unreleased.md`](docs/changelog/unreleased.md) as you go.
- **Cutting a release**: rename `unreleased.md` to `vX.Y.Z.md`, update the table below, create a new blank `unreleased.md` from the template, and add any breaking changes to `docs/breaking-changes.md`.

---

## Versions

| Version | Date | Summary |
|---|---|---|
| [Unreleased](docs/changelog/unreleased.md) | — | Work in progress |
| [v0.1.0](docs/changelog/v0.1.0.md) | 2026-04-23 | Social login, PKCE, secrets encryption at rest, secure logout, token reuse detection, CI/CD, production error mode |
