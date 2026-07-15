---
name: scaff/script/bundle-for-linux
description: >-
  Rule script.bundle.for-linux: scaffold script/bundle/for-linux/main.go
  linux/amd64 cross-compile helper. Triggers: linux bundle, cross compile,
  script.bundle.for-linux.
---

# script/bundle-for-linux — rule `script.bundle.for-linux`

Scaffold a linux/amd64 cross-compile helper.

| Field | Value |
|-------|-------|
| Rule ID | `script.bundle.for-linux` |
| Lint | no |
| Fix | yes |
| Files | `script/bundle/for-linux/main.go` |

## Behavior

- **Fix**: creates the bundle script if missing.
- Sets `GOOS=linux`, `GOARCH=amd64`, `CGO_ENABLED=0` and writes a linux binary
  (default name `app-linux-amd64` in the stub).
- Idempotent when the file already exists.

## CLI

```bash
scaff fix script.bundle.for-linux --dry-run
scaff fix script.bundle.for-linux
go run ./script/bundle/for-linux
go run ./script/bundle/for-linux -o myapp-linux-amd64
```

## Related topics

- `script/build`
- `github/release`
