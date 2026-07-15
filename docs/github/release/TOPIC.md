---
name: scaff/github/release
description: >-
  Rule github/release: scaffold script/github/release and lib helpers for
  GitHub Releases. Triggers: release script, upload release assets,
  github/release.
---

# github/release — rule `github/release`

Scaffold GitHub release helper scripts under `script/github/`.

| Field | Value |
|-------|-------|
| Rule ID | `github/release` |
| Lint | no |
| Fix | yes |
| Files | `script/github/release/main.go`, `script/github/lib/build_release.go` |

## Behavior

- **Fix**: creates release entrypoint and shared build helper when missing.
- Typical flow builds multi-platform artifacts and uploads to a GitHub Release
  using credentials (see `github/upload` for credential docs).
- Supports `--dry-run` on the generated release CLI once scaffolded.

## CLI

```bash
scaff fix github/release --dry-run
scaff fix github/release
go run ./script/github/release --dry-run
go run ./script/github/release
```

## Related topics

- `github/upload` — credentials file and token setup (docs-only)
- `script/github/release-assets` — pack a dir and opt-in `gh` upload/clobber
- `github/testing-workflow`
- `script/bundle-for-linux`
