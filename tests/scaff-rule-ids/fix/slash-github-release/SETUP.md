# Scenario

**Feature**: `scaff fix github/release --dry-run` accepts slash ID

```
# second rule pair: slash form accepted for release scaffold
scaff fix github/release --dry-run -> exit 0
```

## Preconditions

- Temp project with `go.mod` module path `github.com/xhd2015/myapp` (release metadata).
- No pre-existing `script/github/` scaffold required for dry-run accept.
- Rule ID is the slash form `github/release` (not `github.release`).

## Steps

1. Write GitHub-oriented `go.mod`.
2. Run `scaff fix github/release --dry-run`.

```go
import (
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "github/release", "--dry-run"}
	return nil
}
```
