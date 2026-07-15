# Scenario

**Feature**: script/generate stub includes Proposed behavior sketch

```
# empty project, no script/generate/main.go
scaff fix script/generate -> main.go with usage + Proposed behavior
```

## Preconditions

- `script/generate/main.go` does not exist.
- Project has minimal `go.mod` from parent.

## Steps

1. Run `scaff fix script/generate`.
2. Assert generated `script/generate/main.go` contains `Proposed behavior`.

```go
import (
	"os"
	"path/filepath"
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	rel := "script/generate/main.go"
	if _, err := os.Stat(filepath.Join(req.ProjectDir, rel)); err == nil {
		return os.Remove(filepath.Join(req.ProjectDir, rel))
	}
	req.Args = []string{"fix", "script/generate"}
	return nil
}
```
