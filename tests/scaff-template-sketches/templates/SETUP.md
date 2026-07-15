# Scenario

**Feature**: fix creates Go scaffolds from empty project fixtures

```
# grouping: all leaves start with no target .go file, then fix one rule
empty project -> scaff fix <slash-id> -> generated template
```

## Preconditions

- Root allocated temp dir and built `scaff` binary.
- Target scaffold paths are **absent** until fix runs.
- Baseline fixture is a minimal Go `go.mod` unless a leaf overrides (release).

## Steps

1. Write default `go.mod` for a generic Go project.
2. Leaf setups select the fix rule (and may replace `go.mod` for GitHub metadata).
3. Run `scaff fix` and assert sketch text in the created file.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```
