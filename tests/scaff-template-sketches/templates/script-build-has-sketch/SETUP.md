# Scenario

**Feature**: script/build stub includes Proposed behavior sketch

```
# empty project, no script/build/build.go
scaff fix script/build -> build.go with usage + Proposed behavior
```

## Preconditions

- `script/build/build.go` does not exist.
- Project has minimal `go.mod` from parent.

## Steps

1. Run `scaff fix script/build`.
2. Assert generated `script/build/build.go` contains `Proposed behavior`.

```go
import (
	"os"
	"path/filepath"
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	rel := "script/build/build.go"
	if _, err := os.Stat(filepath.Join(req.ProjectDir, rel)); err == nil {
		return os.Remove(filepath.Join(req.ProjectDir, rel))
	}
	req.Args = []string{"fix", "script/build"}
	return nil
}
```
