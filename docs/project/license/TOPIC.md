---
name: scaff/project/license
description: >-
  Rule project/license: ensure root MIT LICENSE with year and owner metadata.
  Triggers: missing LICENSE, MIT license scaffold.
---

# project/license — rule `project/license`

Ensure a root MIT `LICENSE` file exists.

| Field | Value |
|-------|-------|
| Rule ID | `project/license` |
| Lint | yes (default) |
| Fix | yes |
| Files | `LICENSE` |

## Behavior

- **Lint**: missing `LICENSE` → missing; present → OK (presence only).
- **Fix**: create-if-absent with MIT text using `__YEAR__` and `__OWNER__`
  from project metadata (`go.mod` / directory name).
- Idempotent when `LICENSE` already exists.

## CLI

```bash
scaff lint
scaff fix project/license --dry-run
scaff fix project/license
```

## Related topics

- `project/readme`
- `lint`
