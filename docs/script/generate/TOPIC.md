---
name: scaff/script/generate
description: >-
  Rule script.generate: scaffold script/generate/main.go no-op stub for code
  generators. Triggers: generate script, script.generate, codegen entrypoint.
---

# script/generate — rule `script.generate`

Scaffold a generate entrypoint for project code generators.

| Field | Value |
|-------|-------|
| Rule ID | `script.generate` |
| Lint | no |
| Fix | yes |
| Files | `script/generate/main.go` |

## Behavior

- **Fix**: creates `script/generate/main.go` as a no-op stub if missing.
- Idempotent when the file already exists.
- **Dry-run**: reports that the file would be created.

## CLI

```bash
scaff fix script.generate --dry-run
scaff fix script.generate
go run ./script/generate
```

## Related topics

- `script/install`
- `script/build`
- `fix`
