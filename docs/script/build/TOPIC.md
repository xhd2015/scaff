---
name: scaff/script/build
description: >-
  Rule script.build: scaffold script/build/build.go native go build helper.
  Triggers: build script, go build helper, script.build, bin/app.
---

# script/build — rule `script.build`

Scaffold a native `go build` helper.

| Field | Value |
|-------|-------|
| Rule ID | `script.build` |
| Lint | no |
| Fix | yes |
| Files | `script/build/build.go` |

## Behavior

- **Fix**: creates `script/build/build.go` if missing (builds to `bin/app` by
  default in the stub).
- Idempotent when the file already exists.
- Used by `script.install` and as a local developer entrypoint.

## CLI

```bash
scaff fix script.build --dry-run
scaff fix script.build
go run ./script/build
```

## Related topics

- `script/install`
- `script/bundle-for-linux`
