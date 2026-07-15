---
name: scaff/lint
description: >-
  scaff lint: audit default scaffolding rules on a project. Triggers: audit
  project, check scaffolding, lint gitignore, missing workflow.
---

# lint — audit scaffolding

```bash
scaff lint [--dir DIR] [--json] [--profile PROFILE]
```

## Behavior

- Runs **lint-enabled** rules only (today: `git/ignore`, `github/testing-workflow`).
- Exit code non-zero when any rule reports issues (missing/partial).
- `--json` emits a machine-readable report.
- `--profile` overrides auto-detect (`go`, `node`, `polyglot`, `generic`).
- `--dir` selects project root (default: current directory).

## Example

```bash
scaff lint
scaff lint --dir ./my-app --json
scaff lint --profile go
```

## Rule statuses

| Status | Meaning |
|--------|---------|
| OK | expected scaffolding present |
| partial | some patterns/files present, others missing |
| missing | required file or all expected patterns absent |

## Related topics

- `overview` — product model
- `fix` — apply a single rule
- `git/ignore` — lint rule details
- `github/testing-workflow` — lint rule details

## Related CLI

```bash
scaff rules
scaff fix git/ignore --dry-run
```
