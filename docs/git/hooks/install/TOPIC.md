---
name: scaff/git/hooks/install
description: >-
  Rule git.hooks.install: install scaff hooks into .git/hooks/pre-commit and
  pre-push. Triggers: install git hooks, patch pre-commit, git.hooks.install.
---

# git/hooks/install — rule `git.hooks.install`

Patch repository git hooks so they call the scaff git-hooks runner.

| Field | Value |
|-------|-------|
| Rule ID | `git.hooks.install` |
| Lint | no |
| Fix | yes |
| Files | `.git/hooks/pre-commit`, `.git/hooks/pre-push` |

## Preconditions

- `script/git-hooks/main.go` must exist (run `scaff fix git.hooks` first).
- Project must be a git repository (`git` dir resolvable from project root).

## Behavior

- **Fix**: ensures hook scripts invoke `go run ./script/git-hooks <hook>` with
  the marker `# scaff hooks` (idempotent patch).
- **Dry-run**: reports that pre-commit and pre-push would be patched.
- Fails clearly if the runner script is missing.

## CLI

```bash
scaff fix git.hooks
scaff fix git.hooks.install --dry-run
scaff fix git.hooks.install
```

## Related topics

- `git/hooks` — runner script rule
- `fix`
