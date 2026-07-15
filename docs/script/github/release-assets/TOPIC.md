---
name: scaff/script/github/release-assets
description: >-
  Rule script/github/release-assets: scaffold script/github/release-assets/main.go
  helper to pack a directory of assets and optionally upload via gh.
  Triggers: release assets, gh upload, clobber assets, --upload, --dir.
---

# script/github/release-assets — rule `script/github/release-assets`

Scaffold a generic GitHub release-asset helper under `script/github/release-assets/`.

| Field | Value |
|-------|-------|
| Rule ID | `script/github/release-assets` |
| Lint | no |
| Fix | yes |
| Files | `script/github/release-assets/main.go` |

## Behavior

- **Fix**: creates `script/github/release-assets/main.go` if missing.
- Template includes a **Proposed behavior** sketch and packable `--help` text.
- Generated CLI:
  - `--dir` — directory of assets to pack / upload (required at runtime)
  - `--upload` — **opt-in** publish via `gh` (default is plan-only)
  - `--tag` / `--title` — release metadata
  - upload path uses `gh release upload ... --clobber` to replace same-named assets
- Generic: not tied to a single product binary or agent toolchain.
- Related to, but separate from, `github/release` (build + full release pipeline).

## CLI

```bash
scaff fix script/github/release-assets --dry-run
scaff fix script/github/release-assets
go run ./script/github/release-assets --help
go run ./script/github/release-assets --dir ./dist
go run ./script/github/release-assets --dir ./dist --upload --tag v1.0.0
```

## Related topics

- `github/release` — multi-platform release build + upload scaffold
- `github/upload` — credentials / token ops (docs-only)
- `fix`
