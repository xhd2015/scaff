# Scenario

**Feature**: fix git/hooks is idempotent

```
# existing hook runner -> no-op
git/hooks fix -> nothing to do
```

## Preconditions

- `script/git-hooks/main.go` already exists.

## Steps

1. Write existing hook runner.
2. Run `scaff fix git/hooks`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeGitHooksMain(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/hooks"}
	return nil
}
```