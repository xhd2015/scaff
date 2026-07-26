---
name: scaff/install-via-curl
description: >-
  Rule install/via-curl: scaffold install.sh curl installer at repo root.
  Triggers: curl install script, install.sh, install-via-curl, install/via-curl.
---

# install-via-curl — rule `install/via-curl`

Scaffold a root-level curl installer script for release binaries.

| Field | Value |
|-------|-------|
| Rule ID | `install/via-curl` |
| Lint | no |
| Fix | yes |
| Files | `install.sh` |

## Behavior

- **Fix**: creates `install.sh` if missing.
- If legacy `install-via-curl.sh` exists and `install.sh` does not, **renames**
  it to `install.sh` (no content rewrite).
- Detects OS/arch, resolves latest (or tagged) GitHub release asset, downloads
  and installs the binary.
- Idempotent when `install.sh` already exists.
- **Dry-run**: reports create or rename without writing.

## CLI

```bash
scaff fix install/via-curl --dry-run
scaff fix install/via-curl
# after publishing releases, end users may:
# curl -fsSL .../install.sh | bash
```

## Related topics

- `github/release` — publishing assets the installer downloads
- `script/install` — local `go install` helper for developers
