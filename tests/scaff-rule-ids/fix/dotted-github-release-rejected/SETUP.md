# Scenario

**Feature**: `scaff fix github.release` rejects dotted legacy ID

```
# strict: no alias from github.release to github/release
scaff fix github.release -> unknown rule, exit ≠ 0
```

## Preconditions

- Temp Go project ready (content irrelevant for unknown-rule path).
- Rule argument is the **legacy dotted** ID `github.release`.

## Steps

1. Run `scaff fix github.release`.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "github.release"}
	return nil
}
```
