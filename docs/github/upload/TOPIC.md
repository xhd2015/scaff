---
name: scaff/github/upload
description: >-
  Docs-only topic for GitHub upload credentials and release asset upload ops.
  No scaff fix rule. Triggers: upload credentials, .upload-credentials.json,
  GitHub token for release, github/upload.
---

# github/upload — credentials and upload ops (docs-only)

This topic documents how release scripts authenticate to GitHub. There is
**no** scaff fix rule named `github.upload` and no auto-generated credentials
file from `scaff fix`.

## Credentials file

Release helpers typically load:

```text
.upload-credentials.json
```

Expected shape (illustrative):

```json
{
  "token": "ghp_...",
  "owner": "your-org-or-user",
  "repo": "your-repo"
}
```

- Keep this file **out of git** (add to `.gitignore`).
- Prefer fine-scoped tokens with release/upload permissions only.
- Do not commit tokens or paste them into skill install targets.

## Operational flow

1. Scaffold release scripts: `scaff fix github.release`
2. Place credentials locally (not via scaff)
3. Run `go run ./script/github/release --dry-run` then without dry-run

## CLI (related)

```bash
scaff skill --show github/release
scaff fix github.release
go run ./script/github/release --dry-run
```

## Related topics

- `github/release` — scaffold release scripts (rule `github.release`)
- `git/ignore` — ignore secrets and build artifacts
