# Scenario

**Feature**: github/release main template includes Proposed behavior sketch

```
# empty project with GitHub go.mod, no release scaffold
scaff fix github/release -> release/main.go with Proposed behavior
```

## Preconditions

- No `script/github/release/main.go` (or lib) yet.
- `go.mod` uses module path `github.com/xhd2015/myapp` for metadata substitution.

## Steps

1. Write GitHub-oriented `go.mod` (replaces parent generic module).
2. Run `scaff fix github/release`.
3. Assert `script/github/release/main.go` contains `Proposed behavior`.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "github/release"}
	return nil
}
```
