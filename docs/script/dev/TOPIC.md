---
name: scaff/script/dev
description: >-
  Rule script/dev: scaffold script/dev/main.go go run . --dev wrapper.
  Triggers: dev script, local dev runner, go run --dev.
---

# script/dev — rule `script/dev`

Scaffold a local dev wrapper that runs `go run . --dev`.

| Field | Value |
|-------|-------|
| Rule ID | `script/dev` |
| Lint | no (opt-in fix) |
| Fix | yes |
| Files | `script/dev/main.go` |

## Behavior

- **Fix**: creates `script/dev/main.go` if missing.
- Stub forwards optional args to `go run . --dev`.
- Idempotent when the file already exists.
- Not part of default lint.

## CLI

```bash
scaff fix script/dev --dry-run
scaff fix script/dev
go run ./script/dev
```

## Related topics

- `script/build`
- `script/install`
