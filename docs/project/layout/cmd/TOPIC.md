---
name: scaff/project/layout/cmd
description: >-
  Rule project/layout/cmd: scaffold cmd/<name>/main.go CLI entry from module
  name. Triggers: missing cmd layout, CLI main entry, cmd package.
---

# project/layout/cmd — rule `project/layout/cmd`

Scaffold a `cmd/<name>/main.go` entrypoint when `cmd/` is missing.

| Field | Value |
|-------|-------|
| Rule ID | `project/layout/cmd` |
| Lint | no (opt-in fix) |
| Fix | yes |
| Files | `cmd/<name>/main.go` |

## Behavior

- **Fix**: if `cmd/` already exists → no-op; otherwise create
  `cmd/<meta.Name>/main.go` stub with `run` returning not implemented.
- Name comes from project metadata (`go.mod` last path segment or directory).
- Not part of default lint.

## CLI

```bash
scaff fix project/layout/cmd --dry-run
scaff fix project/layout/cmd
```

## Related topics

- `project/agents`
- `script/dev`
