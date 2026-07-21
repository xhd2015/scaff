# Scenario

**Feature**: fix git/pre-commit is idempotent

```
# existing pre-commit helper -> no-op
git/pre-commit fix -> nothing to do
```

## Preconditions

- `script/git/pre-commit/main.go` already exists.

## Steps

1. Write existing pre-commit helper.
2. Run `scaff fix git/pre-commit`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeGitPreCommitMain(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/pre-commit"}
	return nil
}
```
