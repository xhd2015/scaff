---
name: scaff
description: >-
  Multi-topic skill for scaff: amend scaffolding on existing projects.
  Use when auditing or applying git ignore, GitHub workflows, script helpers,
  hooks, release tooling, or curl installers. Retrieve nested topics via the
  skill show action with a slash-separated topic path.
---

# scaff — multi-topic scaffolding skill

This skill is an **index** over nested topics. Load detailed guidance with
`scaff skill --show <topic>` (or `scaff skill <topic> --show`). Topics are
organized as a tree; address a nested topic with a slash-separated path
(e.g. `git/ignore`, `script/bundle-for-linux`).

**scaff** amends scaffolding on *existing* projects (not greenfield create).
Domain commands remain `scaff lint`, `scaff fix <rule>`, and `scaff rules`.
This skill documents product model, CLI surfaces, and each scaffolding rule.

## Topics

- `overview` — product model: amend-not-create, profiles, dry-run/idempotency
- `lint` — `scaff lint` audit CLI
- `fix` — `scaff fix` apply-one-rule CLI
- `git/ignore` — rule `git/ignore` (.gitignore patterns by profile)
- `git/hooks` — rule `git/hooks` (script/git-hooks runner)
  - `install` — rule `git/hooks/install` (patch `.git/hooks/`)
- `github/testing-workflow` — rule `github/testing-workflow` (CI test workflow)
- `github/release` — rule `github/release` (GitHub Releases helper scripts)
- `github/upload` — credentials / upload ops (docs-only; no fix rule)
- `project/readme` — rule `project/readme` (root README.md)
- `project/license` — rule `project/license` (root MIT LICENSE)
- `project/agents` — rule `project/agents` (root AGENTS.md)
- `project/layout/cmd` — rule `project/layout/cmd` (cmd/<name>/main.go)
- `tests/doctest` — rule `tests/doctest` (tests/<name>-cli doctest harness)
- `script/generate` — rule `script/generate`
- `script/install` — rule `script/install`
- `script/build` — rule `script/build`
- `script/dev` — rule `script/dev` (go run . --dev wrapper)
- `script/bundle-for-linux` — rule `script/bundle/for-linux`
- `script/github/release-assets` — rule `script/github/release-assets` (gh release asset pack/upload)
- `install-via-curl` — rule `install/via-curl` (root curl installer)

## Retrieve topics

```bash
# list skill name + every nested topic path
scaff skill --list

# root skill index (this document)
scaff skill --show

# top-level topic
scaff skill --show overview
scaff skill --show lint
scaff skill --show fix

# nested topic (slash path; both flag orders)
scaff skill --show git/ignore
scaff skill git/ignore --show
scaff skill --show git/hooks/install
scaff skill --show github/release
scaff skill --show github/upload
scaff skill --show project/readme
scaff skill --show project/layout/cmd
scaff skill --show tests/doctest
scaff skill --show script/dev
scaff skill --show script/bundle-for-linux
scaff skill --show script/github/release-assets
scaff skill --show install-via-curl

# YAML frontmatter only
scaff skill --show --header
scaff skill --show git/ignore --header
```

## Related CLI

```bash
scaff lint [--dir DIR] [--json] [--profile PROFILE]
scaff fix <rule> [--dir DIR] [--dry-run]
scaff rules [--json]
scaff skill --help
```
