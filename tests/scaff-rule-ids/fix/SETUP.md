# Scenario

**Feature**: `scaff fix` accepts only slash rule IDs

```
# fix CLI dispatch: slash IDs known, dotted IDs unknown
User -> scaff fix <rule-id> [--dry-run] -> exit / stderr
```

## Preconditions

- A fresh temporary project directory exists for each leaf.
- The `scaff` binary is built from module root (`./cmd/scaff`) once per session.
- Default fixture is a minimal Go project (`go.mod`); release leaves may upgrade
  the module path.

## Steps

1. Allocate `req.ProjectDir` / `req.RunDir`.
2. Build or reuse the session-cached `scaff` binary into `req.ScaffBin`.
3. Write a baseline `go.mod` so fix resolvers can detect a Go profile.
4. Leaf setups set `req.Args` to the rule ID under test (slash or dotted).

## Context

- Exit 0 on successful fix / dry-run for a known slash rule.
- Exit ≠ 0 (product uses 2) for unknown rules, including all legacy dotted IDs.
- No dotted → slash aliases.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	req.ProjectDir = t.TempDir()
	req.RunDir = req.ProjectDir
	req.ScaffBin = buildScaffBinary(t)
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```
