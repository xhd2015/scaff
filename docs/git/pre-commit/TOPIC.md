---
name: scaff/git/pre-commit
description: >-
  Rule git/pre-commit: scaffold script/git/pre-commit for ensure-if-missing
  paths and git add. Install via git-hooks. Triggers: pre-commit script,
  git/pre-commit, ensure placeholder, git-hooks pre-commit add.
---

# git/pre-commit — rule `git/pre-commit`

Scaffold a brief pre-commit helper at `script/git/pre-commit/main.go`.

| Field | Value |
|-------|-------|
| Rule ID | `git/pre-commit` |
| Lint | no |
| Fix | yes |
| Files | `script/git/pre-commit/main.go` |

## Behavior

- **Fix**: creates `script/git/pre-commit/main.go` if missing.
- Stub ensures listed paths exist (empty file if missing), then `git add`s them.
- Silent on success; `Error:` on stderr and non-zero exit on failure.
- Edit the `ensure` slice in the generated file for your repo paths.
- Idempotent when the file already exists.
- Does **not** install hooks. Use `git-hooks` (comment in the file header).

## Install

```bash
git-hooks pre-commit add 'script.git.pre-commit' go run ./script/git/pre-commit
```

## CLI

```bash
scaff fix git/pre-commit --dry-run
scaff fix git/pre-commit
go run ./script/git/pre-commit
```

## Related topics

- `git/hooks`
- `git/hooks/install`
- `fix`
