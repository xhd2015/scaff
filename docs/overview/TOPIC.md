---
name: scaff/overview
description: >-
  Product model for scaff: amend scaffolding on existing projects, profiles,
  lint vs fix, dry-run and idempotency. Triggers: what is scaff, how scaff
  works, project profiles, scaffolding philosophy.
---

# overview — product model

**scaff** amends scaffolding on *existing* repositories. It does not create
new projects from templates (use other tools for greenfield create).

## Core ideas

| Concept | Meaning |
|---------|---------|
| Amend, not create | Detect what is missing/partial and fill gaps |
| Profiles | `go`, `node`, `polyglot`, `generic` — auto-detected or `--profile` |
| Lint | Audit rules; report OK / partial / missing |
| Fix | Apply **one** rule at a time; never a bulk fix-all |
| Dry-run | `--dry-run` on fix shows planned actions without writing |
| Idempotent | Re-running fix when already satisfied is a no-op |

## Commands

```bash
# audit default lint rules
scaff lint [--dir DIR] [--json] [--profile PROFILE]

# apply a single rule
scaff fix <rule> [--dir DIR] [--dry-run]

# list lint + fix rules
scaff rules [--json]
```

Default lint rules:

- `git/ignore`
- `github/testing-workflow`

All other catalog rules are fix-only (or docs-only for `github/upload`).

## Project detection

`scaff` resolves a project root (`--dir` or `.`), detects language profile from
files such as `go.mod` / `package.json`, and runs rules against that context.

## Skill topics

Scaffolding recipes live under this multi-topic skill. Retrieve details with:

```bash
scaff skill --show git/ignore
scaff skill --show github/testing-workflow
scaff skill --list
```

## Related topics

- `lint` — lint CLI surface
- `fix` — fix CLI surface
- `git/ignore`, `github/*`, `script/*`, `install-via-curl` — individual rules
