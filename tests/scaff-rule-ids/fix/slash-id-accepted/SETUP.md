# Scenario

**Feature**: `scaff fix git/ignore --dry-run` accepts slash ID

```
# slash form is the only accepted id for git ignore fix
scaff fix git/ignore --dry-run -> exit 0
```

## Preconditions

- Temp Go project with no `.gitignore` (or incomplete; dry-run must still resolve).
- Rule ID is the slash form `git/ignore` (not `git.ignore`).

## Steps

1. Run `scaff fix git/ignore --dry-run` from the project directory.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "git/ignore", "--dry-run"}
	return nil
}
```
