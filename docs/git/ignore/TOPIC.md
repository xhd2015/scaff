---
name: scaff/git/ignore
description: >-
  Rule git.ignore: ensure .gitignore has common patterns for the project
  profile. Triggers: missing gitignore, add node_modules, ignore bin/,
  git.ignore fix.
---

# git/ignore — rule `git.ignore`

Ensure `.gitignore` includes profile-appropriate ignore patterns.

| Field | Value |
|-------|-------|
| Rule ID | `git.ignore` |
| Lint | yes |
| Fix | yes |
| Files | `.gitignore` |

## Behavior

- **Lint**: compares existing `.gitignore` to expected patterns for the project
  profile (`go`, `node`, `polyglot`, `generic`). Reports missing patterns.
- **Fix**: appends only missing patterns (idempotent). Does not rewrite the
  whole file.
- **Dry-run**: `scaff fix git.ignore --dry-run` lists patterns that would be
  appended without writing.

## CLI

```bash
scaff lint
scaff fix git.ignore --dry-run
scaff fix git.ignore
scaff fix git.ignore --dir ./my-app
```

## Idempotency

Re-running fix when all patterns are present prints that nothing is needed and
does not change the file.

## Related topics

- `overview` — profiles and amend model
- `lint` — audit surface
- `fix` — apply-one-rule surface
