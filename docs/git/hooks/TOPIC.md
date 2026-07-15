---
name: scaff/git/hooks
description: >-
  Rule git.hooks: add script/git-hooks runner for install, pre-commit, and
  pre-push. Triggers: git hooks script, pre-commit scaffolding, git.hooks.
---

# git/hooks — rule `git.hooks`

Scaffold a `script/git-hooks` runner used by local git hooks.

| Field | Value |
|-------|-------|
| Rule ID | `git.hooks` |
| Lint | no |
| Fix | yes |
| Files | `script/git-hooks/main.go` |

## Behavior

- **Fix**: creates `script/git-hooks/main.go` if missing (stub supporting
  `install`, `pre-commit`, `pre-push`).
- Idempotent when the file already exists.
- Does **not** by itself patch `.git/hooks/` — use `git.hooks.install` after
  the runner exists.

## CLI

```bash
scaff fix git.hooks --dry-run
scaff fix git.hooks
go run ./script/git-hooks install
go run ./script/git-hooks pre-commit
go run ./script/git-hooks pre-push
```

## Nested topic

- `git/hooks/install` — rule `git.hooks.install` patches `.git/hooks/pre-commit`
  and `pre-push` to invoke this runner.

## Related topics

- `git/hooks/install`
- `fix`
