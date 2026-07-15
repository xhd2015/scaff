---
name: scaff/fix
description: >-
  scaff fix: apply one scaffolding rule at a time. Triggers: fix gitignore,
  add workflow, scaffold script, dry-run fix, apply rule.
---

# fix — apply one rule

```bash
scaff fix <rule> [--dir DIR] [--dry-run]
```

## Behavior

- Exactly **one** rule per invocation (no fix-all).
- Unknown rule → non-zero exit and list of available fix rules.
- `--dry-run` prints planned actions without writing files.
- Idempotent: if the target already exists / is complete, reports nothing to do.

## Examples

```bash
scaff fix git/ignore --dry-run
scaff fix git/ignore
scaff fix github/testing-workflow --dir ./svc
scaff fix script/build
scaff fix script/github/release-assets
scaff fix install/via-curl
```

## Fix rule IDs

| Rule ID | Topic path |
|---------|------------|
| `git/ignore` | `git/ignore` |
| `git/hooks` | `git/hooks` |
| `git/hooks/install` | `git/hooks/install` |
| `github/testing-workflow` | `github/testing-workflow` |
| `github/release` | `github/release` |
| `script/generate` | `script/generate` |
| `script/install` | `script/install` |
| `script/build` | `script/build` |
| `script/bundle/for-linux` | `script/bundle-for-linux` |
| `script/github/release-assets` | `script/github/release-assets` |
| `install/via-curl` | `install-via-curl` |

There is **no** fix rule for GitHub upload credentials (`github/upload` is docs-only).

## Related topics

- `overview` — product model and idempotency
- `lint` — audit before fix
- individual rule topics under `git/`, `github/`, `script/`, `install-via-curl`
