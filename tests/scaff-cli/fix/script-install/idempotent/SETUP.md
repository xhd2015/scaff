# Scenario

**Feature**: fix script/install is idempotent

```
# existing install stub -> no-op
script/install fix -> nothing to do
```

## Preconditions

- `script/install/install.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script/install`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeScriptInstall(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/install"}
	return nil
}
```