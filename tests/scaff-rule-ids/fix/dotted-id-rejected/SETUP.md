# Scenario

**Feature**: `scaff fix git.ignore` rejects dotted legacy ID

```
# strict: no alias from git.ignore to git/ignore
scaff fix git.ignore -> unknown rule, exit ≠ 0
```

## Preconditions

- Temp Go project ready (content irrelevant for unknown-rule path).
- Rule argument is the **legacy dotted** ID `git.ignore`.

## Steps

1. Run `scaff fix git.ignore` (no dry-run required; must fail before apply).

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "git.ignore"}
	return nil
}
```
