---
name: scaff/script/install
description: >-
  Rule script.install: scaffold script/install/install.go build-then-install
  helper. Triggers: install script, go install helper, script.install.
---

# script/install — rule `script.install`

Scaffold a build-then-install helper.

| Field | Value |
|-------|-------|
| Rule ID | `script.install` |
| Lint | no |
| Fix | yes |
| Files | `script/install/install.go` |

## Behavior

- **Fix**: creates `script/install/install.go` if missing.
- Typical flow: run `./script/build`, then `go install .`.
- Idempotent when the file already exists.

## CLI

```bash
scaff fix script.install --dry-run
scaff fix script.install
go run ./script/install
```

## Related topics

- `script/build`
- `script/generate`
- `install-via-curl`
