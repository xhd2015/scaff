---
name: scaff/project/readme
description: >-
  Rule project/readme: ensure root README.md with profile-aware install notes.
  Triggers: missing README, project readme, install instructions.
---

# project/readme — rule `project/readme`

Ensure a root `README.md` exists for the project.

| Field | Value |
|-------|-------|
| Rule ID | `project/readme` |
| Lint | yes (default) |
| Fix | yes |
| Files | `README.md` |

## Behavior

- **Lint**: missing `README.md` → missing; present → OK.
- **Fix**: create-if-absent only (idempotent when the file already exists).
- Template includes `# __NAME__`, profile-specific Install section, and Usage.
  - Go: `go install __MODULE__@latest`
  - Node: `npm install` / `npm run dev`
  - Generic: no Install section

## CLI

```bash
scaff lint
scaff fix project/readme --dry-run
scaff fix project/readme
```

## Related topics

- `project/license`
- `project/agents`
- `lint`
