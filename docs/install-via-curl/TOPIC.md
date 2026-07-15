---
name: scaff/install-via-curl
description: >-
  Rule install.via.curl: scaffold install-via-curl.sh curl installer at repo
  root. Triggers: curl install script, install-via-curl, install.via.curl.
---

# install-via-curl — rule `install.via.curl`

Scaffold a root-level curl installer script for release binaries.

| Field | Value |
|-------|-------|
| Rule ID | `install.via.curl` |
| Lint | no |
| Fix | yes |
| Files | `install-via-curl.sh` |

## Behavior

- **Fix**: creates `install-via-curl.sh` if missing.
- Detects OS/arch, resolves latest (or tagged) GitHub release asset, downloads
  and installs the binary.
- Idempotent when the script already exists.
- **Dry-run**: reports that the script would be created.

## CLI

```bash
scaff fix install.via.curl --dry-run
scaff fix install.via.curl
# after publishing releases, end users may:
# curl -fsSL .../install-via-curl.sh | bash
```

## Related topics

- `github/release` — publishing assets the installer downloads
- `script/install` — local `go install` helper for developers
