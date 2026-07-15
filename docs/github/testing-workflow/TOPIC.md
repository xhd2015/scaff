---
name: scaff/github/testing-workflow
description: >-
  Rule github.testing.workflow: ensure .github/workflows/test.yml runs go test
  and doctest. Triggers: missing CI, add test workflow, github.testing.workflow.
---

# github/testing-workflow — rule `github.testing.workflow`

Ensure a GitHub Actions test workflow exists for the project.

| Field | Value |
|-------|-------|
| Rule ID | `github.testing.workflow` |
| Lint | yes |
| Fix | yes |
| Files | `.github/workflows/test.yml` |

## Behavior

- **Lint**: checks that `.github/workflows/test.yml` exists.
- **Fix**: creates the workflow from an embedded template when missing
  (includes `go test` and doctest steps when applicable). Does not overwrite
  an existing workflow file.
- **Dry-run**: reports that the workflow would be created.

## CLI

```bash
scaff lint
scaff fix github.testing.workflow --dry-run
scaff fix github.testing.workflow
```

## Idempotency

If the workflow file already exists, fix reports nothing to do.

## Related topics

- `lint`
- `github/release`
- `github/upload`
