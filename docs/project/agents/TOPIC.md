---
name: scaff/project/agents
description: >-
  Rule project/agents: scaffold root AGENTS.md with build and test notes.
  Triggers: AGENTS.md, agent instructions, project agents file.
---

# project/agents — rule `project/agents`

Scaffold a root `AGENTS.md` for agent/tooling guidance.

| Field | Value |
|-------|-------|
| Rule ID | `project/agents` |
| Lint | no (opt-in fix) |
| Fix | yes |
| Files | `AGENTS.md` |

## Behavior

- **Fix**: create-if-absent with overview, build (`go run ./script/build`), and
  test (`go test` + doctest) sections using `__NAME__`.
- Idempotent when `AGENTS.md` already exists.
- Not part of default lint.

## CLI

```bash
scaff fix project/agents --dry-run
scaff fix project/agents
```

## Related topics

- `project/readme`
- `tests/doctest`
- `script/build`
