---
name: scaff/tests/doctest
description: >-
  Rule tests/doctest: ensure tests/<name>-cli DOCTEST.md + SETUP.md harness.
  Triggers: missing doctest, CLI tests scaffold, doctest tree.
---

# tests/doctest — rule `tests/doctest`

Ensure a doctest harness under `tests/<name>-cli/`.

| Field | Value |
|-------|-------|
| Rule ID | `tests/doctest` |
| Lint | yes (default, Go profile) |
| Fix | yes |
| Files | `tests/<name>-cli/DOCTEST.md`, `tests/<name>-cli/SETUP.md` |

## Behavior

- **Lint**: Go profile requires `tests/<name>-cli/DOCTEST.md`; node/generic → n/a OK.
- **Fix**: create-if-absent for `DOCTEST.md` + `SETUP.md` (if DOCTEST already
  exists, fix is a no-op even when SETUP is missing).
- Name comes from project metadata.

## CLI

```bash
scaff lint
scaff fix tests/doctest --dry-run
scaff fix tests/doctest
```

## Related topics

- `project/agents`
- `github/testing-workflow`
- `lint`
